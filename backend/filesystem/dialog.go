package filesystem

import (
	"context"
	"errors"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"io/fs"
	"os/exec"
	"path/filepath"
	goRuntime "runtime"
	"strings"
)

type Dialog struct {
	ctx context.Context
}

func NewDialog() *Dialog {
	return &Dialog{}
}

func (d *Dialog) SetContext(ctx context.Context) {
	d.ctx = ctx
}

func (d *Dialog) OpenMultipleFilesDialog() ([]string, error) {
	return runtime.OpenMultipleFilesDialog(d.ctx, runtime.OpenDialogOptions{
		Title: "Select file(s)",
		Filters: []runtime.FileFilter{
			{DisplayName: "JavaScript", Pattern: "*.js"},
			{DisplayName: "PHP", Pattern: "*.php"},
			{DisplayName: "All Files", Pattern: "*"},
		},
		ShowHiddenFiles: true,
	})
}

func (d *Dialog) OpenDirectoryDialog() ([]string, error) {
	path, err := runtime.OpenDirectoryDialog(d.ctx, runtime.OpenDialogOptions{
		Title: "Select folder",
	})
	if err != nil {
		return nil, err
	}

	if path == "" {
		return nil, errors.New("could not open folder")
	}

	var files []string
	err = filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && (strings.HasSuffix(path, ".js") || strings.HasSuffix(path, ".php")) {
			files = append(files, path)
		}

		return nil
	})
	return files, err
}

func (d *Dialog) OpenFileLocation(path string) error {
	currOS := goRuntime.GOOS
	command, ok := map[string]string{
		"windows": "explorer",
		"darwin":  "open",
		"linux":   "xdg-open",
	}[currOS]
	if !ok {
		return errors.New("unsupported OS: " + currOS)
	}

	return exec.Command(command, filepath.Dir(path)).Run()
}

func (d *Dialog) MessageInfoDialog(title string, message string) error {
	_, err := runtime.MessageDialog(d.ctx, runtime.MessageDialogOptions{
		Title:   title,
		Message: message,
		Type:    runtime.InfoDialog,
	})

	return err
}

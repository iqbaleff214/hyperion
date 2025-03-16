package filesystem

import (
	"context"
	"errors"
	"fmt"
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
	paths, err := runtime.OpenMultipleFilesDialog(d.ctx, runtime.OpenDialogOptions{
		Title: "Select file(s)",
		Filters: []runtime.FileFilter{
			{DisplayName: "JavaScript", Pattern: "*.js"},
			{DisplayName: "PHP", Pattern: "*.php"},
			{DisplayName: "Ruby", Pattern: "*.rb"},
			{DisplayName: "Python", Pattern: "*.py"},
			{DisplayName: "All Files", Pattern: "*"},
		},
		ShowHiddenFiles: true,
	})

	fmt.Println(paths)

	return paths, err
}

func (d *Dialog) OpenDirectoryDialog() ([]string, error) {
	path, err := runtime.OpenDirectoryDialog(d.ctx, runtime.OpenDialogOptions{
		Title: "Select folder",
	})
	if err != nil {
		return nil, err
	}

	if path == "" {
		return []string{}, nil
	}

	skipDirs := map[string]bool{
		".git": true,
		// JavaScript / Node.js
		"node_modules": true,
		// PHP
		"vendor": true,
		// Ruby
		"bundle": true,
		".gem":   true,
		// Python
		"__pycache__":   true,
		"env":           true,
		"venv":          true,
		".venv":         true,
		"site-packages": true,
	}

	var files []string
	err = filepath.WalkDir(path, func(path string, dir fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if dir.IsDir() && skipDirs[dir.Name()] {
			return fs.SkipDir
		}

		if !dir.IsDir() && d.hasAllowedSuffix(path) {
			files = append(files, path)
		}

		return nil
	})
	return files, err
}

func (d *Dialog) hasAllowedSuffix(path string) bool {
	allowed := map[string]bool{
		".js":  true,
		".php": true,
		".rb":  true,
		".py":  true,
	}

	for ext := range allowed {
		if strings.HasSuffix(path, ext) {
			return true
		}
	}
	return false
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

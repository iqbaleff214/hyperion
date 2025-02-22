package main

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) OpenFiles() ([]string, error) {
	selection, err := runtime.OpenMultipleFilesDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select files",
		Filters: []runtime.FileFilter{
			{DisplayName: "JavaScript & PHP Files", Pattern: "*.js;*.php"},
			{DisplayName: "All Files", Pattern: "*"},
		},
	})
	if err != nil {
		return nil, err
	}
	return selection, nil
}

func (a *App) OpenFolder() ([]string, error) {
	folder, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select folder",
	})
	if err != nil {
		return nil, err
	}

	if folder == "" {
		return nil, nil
	}

	var files []string
	err = filepath.WalkDir(folder, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && (strings.HasSuffix(path, ".js") || strings.HasSuffix(path, ".php")) {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return files, nil
}

func (a *App) ReadFilesContent(filePaths []string) (map[string]string, error) {
	fileContents := make(map[string]string)

	for _, filePath := range filePaths {
		content, err := os.ReadFile(filePath)
		if err != nil {
			return nil, err
		}
		fileContents[filePath] = string(content)
	}

	return fileContents, nil
}

func (a *App) ObfuscateJS(code string) string {
	obfuscator := NewObfuscator()
	return obfuscator.Obfuscate(code)
}

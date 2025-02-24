package main

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	goRuntime "runtime"

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

func (a *App) ExportEncryptedFiles(files map[string]string) error {
	directoryPath, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select Export Folder",
	})
	if err != nil || directoryPath == "" {
		return err
	}

	for originalPath, encryptedData := range files {
		filename := filepath.Base(originalPath)
		savePath := filepath.Join(directoryPath, filename)
		if _, err := os.Stat(savePath); err == nil {
			ext := filepath.Ext(filename)
			name := filename[:len(filename)-len(ext)]
			count := 1
			for {
				newFilename := fmt.Sprintf("%s_%d%s", name, count, ext)
				newSavePath := filepath.Join(directoryPath, newFilename)
				if _, err := os.Stat(newSavePath); os.IsNotExist(err) {
					savePath = newSavePath
					break
				}
				count++
			}
		}
		err := os.WriteFile(savePath, []byte(encryptedData), 0644)
		if err != nil {
			return err
		}
	}
	runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:    runtime.InfoDialog,
		Title:   "Export Completed",
		Message: "All encrypted files have been successfully exported.",
	})

	return nil
}
func (a *App) OpenImportedFileLocation(filePath string) error {
	dir := filepath.Dir(filePath)
	var cmd *exec.Cmd

	switch goRuntime.GOOS {
	case "windows":
		cmd = exec.Command("explorer", dir)
	case "darwin":
		cmd = exec.Command("open", dir)
	case "linux":
		cmd = exec.Command("xdg-open", dir)
	default:
		return fmt.Errorf("unsupported OS: %s", goRuntime.GOOS)
	}

	return cmd.Start()
}

func (a *App) ObfuscateJS(code string) string {
	obfuscator := NewObfuscator()
	return obfuscator.Obfuscate(code)
}

package filesystem

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
)

type FileManager struct {
	ctx context.Context
}

func NewFileManager() *FileManager {
	return &FileManager{}
}

func (f *FileManager) SetContext(ctx context.Context) {
	f.ctx = ctx
}

func (f *FileManager) ReadFile(path string) (string, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(contents), nil
}

func (f *FileManager) WriteFile(path string, content string) error {
	return os.WriteFile(path, []byte(content), 0644)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err = os.WriteFile(path, []byte(content), 0644); err != nil {
			return err
		}

		return nil
	}

	extension := filepath.Ext(path)
	filename := path[:len(path)-len(extension)]
	index := 1
	for {
		path = fmt.Sprintf("%s_%d%s", filename, index, extension)
		if _, ok := os.Stat(path); os.IsNotExist(ok) {
			if err := os.WriteFile(path, []byte(content), 0644); err != nil {
				return err
			}

			return nil
		}
		index++
	}
}

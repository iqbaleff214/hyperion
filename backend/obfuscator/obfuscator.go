package obfuscator

import (
	"context"
	"errors"
	"log"
	"path/filepath"
)

type Obfuscator interface {
	SetContext(ctx context.Context)
	Config() Config
	Obfuscate(path string) (string, error)
}

type obfuscator struct {
	ctx    context.Context
	config Config
}

func NewObfuscator() Obfuscator {
	return &obfuscator{
		config: NewConfig(),
	}
}

func (o *obfuscator) SetContext(ctx context.Context) {
	o.ctx = ctx
}

func (o *obfuscator) Obfuscate(path string) (string, error) {
	if path == "" {
		return "", errors.New("empty path")
	}

	extension := filepath.Ext(path)
	log.Println(extension)

	return "", nil
}

func (o *obfuscator) Config() Config {
	return o.config
}

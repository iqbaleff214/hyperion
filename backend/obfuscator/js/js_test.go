package js_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"hyperion/backend/obfuscator/js"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var obfuscator = js.Obfuscation{}
var code string

func init() {
	relativePath := "../examples/example.js"
	absolutePath, err := filepath.Abs(relativePath)
	if err != nil {
		panic(err)
	}

	b, err := os.ReadFile(absolutePath)
	if err != nil {
		panic(err)
	}

	code = string(b)
}

func TestObfuscation_RemoveComment(t *testing.T) {
	content, err := obfuscator.RemoveComment(code)
	if err != nil {
		t.Fatal(err)
	}

	assert.Zero(t, strings.Count(content, "//"), "Expected code to not have single-line comments")
	assert.Zero(t, strings.Count(content, "/**"), "Expected code to not have multi-line comments")
}

func TestObfuscation_RenameVariable(t *testing.T) {
	content, err := obfuscator.RenameVariable(code)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(content)
}

func TestObfuscation_RenameConstant(t *testing.T) {
	content, err := obfuscator.RenameConstant(code)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(content)
}

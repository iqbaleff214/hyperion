package php_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"hyperion/backend/obfuscator/php"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var obfuscator = php.Obfuscation{}
var code string

func init() {
	relativePath := "../examples/example.php"
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
	assert.Zero(t, strings.Count(content, "#"), "Expected code to not have multi-line comments")
}

func TestObfuscation_RenameVariable(t *testing.T) {
	_, err := obfuscator.RenameVariable(code)
	if err != nil {
		t.Fatal(err)
	}
}

func TestObfuscation_RenameConstant(t *testing.T) {
	_, err := obfuscator.RenameConstant(code)
	if err != nil {
		t.Fatal(err)
	}
}

func TestObfuscation_UnicodeString(t *testing.T) {
	content, err := obfuscator.UnicodeString(code)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(content)
}

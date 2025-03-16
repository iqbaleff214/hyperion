package php

import (
	"fmt"
	"hyperion/backend/util"
	"regexp"
	"strings"
)

type Obfuscation struct{}

// Minify - minify duh
func (o Obfuscation) Minify(code string) (string, error) {
	var err error
	code, err = o.RemoveComment(code)
	if err != nil {
		return code, err
	}

	code = strings.TrimSpace(code)
	code = regexp.MustCompile(`\s+`).ReplaceAllString(code, " ")

	return code, nil
}

// UnicodeString - string to unicode
func (o Obfuscation) UnicodeString(code string) (string, error) {
	rgx := regexp.MustCompile(`"(?:\\.|[^"])*?"|'(?:\\.|[^'])*?'|` + "`" + `(?:\\.|[^` + "`" + `])*?` + "`")
	code = rgx.ReplaceAllStringFunc(code, func(match string) string {
		quote := match[:1]
		original := match[1 : len(match)-1]
		obfuscated := util.StringToHex(original)
		return fmt.Sprintf("%s%s%s", quote, obfuscated, quote)
	})

	return code, nil
}

func (o Obfuscation) RefactorLoopStatement(code string) (string, error) {
	return code, nil
}

func (o Obfuscation) RefactorIfStatement(code string) (string, error) {
	return code, nil
}

func (o Obfuscation) RenameVariable(code string) (string, error) {
	return code, nil
}

func (o Obfuscation) RenameConstant(code string) (string, error) {
	return code, nil
}

func (o Obfuscation) RenameFunction(code string) (string, error) {
	return code, nil
}

// RemoveComment - remove any comments from code
func (o Obfuscation) RemoveComment(code string) (string, error) {
	singleLine := regexp.MustCompile(`(?m)//.*|#.*`)
	code = singleLine.ReplaceAllString(code, "")

	multiLine := regexp.MustCompile(`(?s)/\*.*?\*/`)
	code = multiLine.ReplaceAllString(code, "")

	return code, nil
}

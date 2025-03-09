package js

import (
	"fmt"
	"hyperion/backend/util"
	"regexp"
	"strings"
)

type Obfuscation struct{}

// SingleLine
// TODO: self-explanatory
func (o Obfuscation) SingleLine(code string) (string, error) {
	return code, nil
}

// RefactorLoopStatement
// TODO: self-explanatory
func (o Obfuscation) RefactorLoopStatement(code string) (string, error) {
	return code, nil
}

// RefactorIfStatement
// TODO: self-explanatory
func (o Obfuscation) RefactorIfStatement(code string) (string, error) {
	return code, nil
}

// RenameVariable
// TODO:
// - Jangan sampai string yang memiliki kata yang sama dengan variable berubah
// - String yang memuat kata yang sama dengan variable bisa terhindar dari perubahan, akan tetapi malah semua variable tidak berubah namanya
func (o Obfuscation) RenameVariable(code string) (string, error) {
	varRegex := regexp.MustCompile(`\b(?P<keyword>let|var)\s+(?P<names>[a-zA-Z_]\w*(?:\s*,\s*[a-zA-Z_]\w*)*)`)
	stringRegex := regexp.MustCompile(`"(?:\\.|[^"])*?"|'(?:\\.|[^'])*?'|` + "`" + `(?:\\.|[^` + "`" + `])*?` + "`")
	templateRegex := regexp.MustCompile(`\$\{([a-zA-Z_]\w*)\}`)

	varMapping := make(map[string]string)
	code = varRegex.ReplaceAllStringFunc(code, func(match string) string {
		parts := varRegex.FindStringSubmatch(match)
		keyword, varNames := parts[1], parts[2]

		names := strings.Split(varNames, ",")
		for i, name := range names {
			name = strings.TrimSpace(name)
			if _, exists := varMapping[name]; !exists {
				varMapping[name] = util.GenerateName("_v")
			}
			names[i] = varMapping[name]
		}

		return keyword + " " + strings.Join(names, ", ")
	})

	matches := stringRegex.FindAllStringIndex(code, -1)

	code = templateRegex.ReplaceAllStringFunc(code, func(match string) string {
		groups := templateRegex.FindStringSubmatch(match)
		oldName := groups[1]

		if newName, exists := varMapping[oldName]; exists {
			return fmt.Sprintf("${%s}", newName)
		}
		return match
	})

	for oldName, newName := range varMapping {
		code = regexp.MustCompile(`\b`+oldName+`\b`).ReplaceAllStringFunc(code, func(match string) string {
			for _, pos := range matches {
				if pos[0] <= strings.Index(code, match) && strings.Index(code, match) <= pos[1] {
					return match
				}
			}
			return newName
		})
	}

	return code, nil
}

// RenameConstant
// TODO:
func (o Obfuscation) RenameConstant(code string) (string, error) {
	rgx := regexp.MustCompile(`(?m)(const\s+)([a-zA-Z_][a-zA-Z0-9_]*)(\s*=)`)
	matches := rgx.FindAllStringSubmatch(code, -1)

	replacements := make(map[string]string)

	for _, match := range matches {
		oldName := match[2]
		if _, exists := replacements[oldName]; !exists {
			replacements[oldName] = util.GenerateName("_c")
		}
	}

	for oldName, newName := range replacements {
		code = strings.ReplaceAll(code, oldName, newName)
	}

	return code, nil
}

// RenameFunction
// TODO: self-explanatory
func (o Obfuscation) RenameFunction(code string) (string, error) {
	return code, nil
}

// RemoveComment - remove any comments from code
func (o Obfuscation) RemoveComment(code string) (string, error) {
	singleLine := regexp.MustCompile(`//.*`)
	code = singleLine.ReplaceAllString(code, "")

	multiLine := regexp.MustCompile(`/\*[\s\S]*?\*/`)
	code = multiLine.ReplaceAllString(code, "")

	return code, nil
}

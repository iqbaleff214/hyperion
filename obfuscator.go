package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

type Obfuscator struct {
	rng         *rand.Rand
	variableMap map[string]string
}

func NewObfuscator() *Obfuscator {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	return &Obfuscator{
		rng:         rng,
		variableMap: make(map[string]string),
	}
}

func (o *Obfuscator) randomVarName() string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	name := fmt.Sprintf("_v%c%d", letters[o.rng.Intn(len(letters))], o.rng.Intn(1000))
	return name
}

func (o *Obfuscator) Obfuscate(jsCode string) string {
	// Obfuscate variable names
	varRegex := regexp.MustCompile(`\b(var|let|const)\s+([a-zA-Z_$][a-zA-Z0-9_$]*)\b`)
	jsCode = varRegex.ReplaceAllStringFunc(jsCode, func(match string) string {
		parts := varRegex.FindStringSubmatch(match)
		keyword, varName := parts[1], parts[2]

		if _, exists := o.variableMap[varName]; !exists {
			o.variableMap[varName] = o.randomVarName()
		}

		return fmt.Sprintf("%s %s", keyword, o.variableMap[varName])
	})

	// Obfuscate normal strings (single and double quotes)
	stringRegex := regexp.MustCompile(`"([^"]*)"|'([^']*)'`)
	jsCode = stringRegex.ReplaceAllStringFunc(jsCode, func(match string) string {
		quote := match[:1]                 // Extract opening quote
		content := match[1:]               // Remove opening quote
		content = content[:len(content)-1] // Remove closing quote

		var encoded strings.Builder
		for _, char := range content {
			encoded.WriteString(fmt.Sprintf("\\x%02x", char))
		}

		return quote + encoded.String() + quote
	})

	// Obfuscate template literals (preserve `${}` expressions and replace variables)
	templateRegex := regexp.MustCompile("`([^`]*)`")
	jsCode = templateRegex.ReplaceAllStringFunc(jsCode, func(match string) string {
		content := match[1 : len(match)-1] // Remove backticks

		var encoded strings.Builder
		inExpression := false
		nestedLevel := 0 // Track `${}` depth
		var expressionBuffer strings.Builder

		for i := 0; i < len(content); i++ {
			if content[i] == '$' && i+1 < len(content) && content[i+1] == '{' {
				inExpression = true
				nestedLevel++
				encoded.WriteString("${")
				i++ // Skip next '{'
				expressionBuffer.Reset()
				continue
			}
			if inExpression {
				if content[i] == '}' {
					nestedLevel--
					if nestedLevel == 0 {
						inExpression = false
						expression := expressionBuffer.String()
						if obfVar, exists := o.variableMap[expression]; exists {
							encoded.WriteString(obfVar)
						} else {
							encoded.WriteString(expression) // Keep unchanged if not found
						}
						encoded.WriteString("}")
						continue
					}
				}
				expressionBuffer.WriteByte(content[i])
			} else {
				encoded.WriteString(fmt.Sprintf("\\x%02x", content[i])) // Encode static text
			}
		}

		return "`" + encoded.String() + "`"
	})
	// Replace variable names inside the entire JS code
	for original, obfuscated := range o.variableMap {
		jsCode = strings.ReplaceAll(jsCode, original, obfuscated)
	}

	// Remove single-line and multi-line comments
	commentRegex := regexp.MustCompile(`(?s)/\*.*?\*/|(?m)^\s*//.*?$`)
	jsCode = commentRegex.ReplaceAllString(jsCode, "")

	// Remove empty lines
	emptyLineRegex := regexp.MustCompile(`(?m)^\s*\n`)
	jsCode = emptyLineRegex.ReplaceAllString(jsCode, "")

	// Ensure each line ends with a semicolon if it doesn't already
	semicolonRegex := regexp.MustCompile(`(?m)([^\s;])\s*$`)
	jsCode = semicolonRegex.ReplaceAllString(jsCode, "$1;")

	// Remove all unnecessary whitespace
	jsCode = regexp.MustCompile(`\s+`).ReplaceAllString(jsCode, " ")
	jsCode = regexp.MustCompile(`\s*([=+\-*/{}(),;])\s*`).ReplaceAllString(jsCode, "$1")

	return jsCode

}

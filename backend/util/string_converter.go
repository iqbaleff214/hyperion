package util

import (
	"fmt"
	"strings"
)

func StringToUnicode(s string) string {
	var result strings.Builder
	for _, r := range s {
		result.WriteString(fmt.Sprintf("\\u%04x", r))
	}
	return result.String()
}

func StringToHex(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] == '\\' && i+1 < len(s) {
			switch s[i+1] {
			case 'n', 't', 'r', '\\', '"', '\'':
				result.WriteByte('\\')
				result.WriteByte(s[i+1])
				i++
				continue
			}
		}
		result.WriteString(fmt.Sprintf("\\x%02X", s[i]))
	}
	return result.String()
}

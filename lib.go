// Conjunto de funções para validação e formatação de documentos brasileiros
package brdocs

import (
	"unicode"
)

func isNumber(s string) bool {
	if s == "" {
		return false
	}

	for _, d := range s {
		if !unicode.IsDigit(d) {
			return false
		}
	}

	return true
}

func getNumbers(s string) string {
	var output string
	for _, d := range s {
		if unicode.IsDigit(d) {
			output += string(d)
		}
	}
	return output
}

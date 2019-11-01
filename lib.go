// Conjunto de funções para validação e formatação de documentos brasileiros
package gobr

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

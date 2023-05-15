package validar

import (
	"regexp"
	"strconv"
)

// CEP valida o CEP informado
func CEP(cep string) bool {
	// Remove caracteres não numéricos
	cep = regexp.MustCompile(`\D`).ReplaceAllString(cep, "")
	// Verifica se o CEP possui 8 dígitos numéricos
	if _, err := strconv.Atoi(cep); err != nil || len(cep) != 8 {
		return false
	}
	return true
}

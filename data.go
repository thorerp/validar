package validar

import (
	"time"
)

// DataPt valida a data informada no formato dd/mm/aaaa
func DataPt(data string) bool {
	// Verifica se o Parse do time consegue interpretar a data, caso não consiga, retorna false
	if _, err := time.Parse("02/01/2006", data); err != nil {
		return false
	}
	return true
}

// DataEn valida a data informada no formato aaaa-mm-dd
func DataEn(data string) bool {
	// Verifica se o Parse do time consegue interpretar a data, caso não consiga, retorna false
	if _, err := time.Parse("2006-01-02", data); err != nil {
		return false
	}
	return true
}

package validar

import "testing"

// ValidaCEPTest a função ValidaCEP
func TestValidaCEP(t *testing.T) {
	casos := []struct {
		cep      string
		validade bool
	}{
		// 5 CEPs diferentes válidos de estados diferentes
		{"12345-678", true},
		{"12345678", true},
		{"12345-678", true},
		{"84060656", true},
		{"84060-144", true},
		// 5 CEPs diferentes inválidos
		{"ABCDE-FGH", false},
		{"ABCDEFGH", false},
		{"12345-67A", false},
		{"Z4060657", false},
		{"1-145", false},
	}
	for _, c := range casos {
		if CEP(c.cep) != c.validade {
			t.Errorf("CEP %s deveria ser %v", c.cep, c.validade)
		}
	}
}

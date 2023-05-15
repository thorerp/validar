package validar

import "testing"

// TestDataPt valida a data informada no formato dd/mm/aaaa
func TestDataPt(t *testing.T) {
	casos := []struct {
		data     string
		validade bool
	}{
		// 5 datas diferentes válidas
		{"01/01/2000", true},
		{"31/12/2000", true},
		{"29/02/2000", true},
		{"29/02/2004", true},
		{"30/04/2020", true},
		// 5 datas diferentes inválidas
		{"01/01/200", false},
		{"00/00/0000", false},
		{"31/11/2000", false},
		{"29/02/2001", false},
		{"29/02/2023", false},
	}
	for _, c := range casos {
		if DataPt(c.data) != c.validade {
			t.Errorf("DataPt %s deveria ser %v", c.data, c.validade)
		}
	}
}

// TestDataEn valida a data informada no formato aaaa-mm-dd
func TestDataEn(t *testing.T) {
	casos := []struct {
		data     string
		validade bool
	}{
		// 5 datas diferentes válidas
		{"2000-01-01", true},
		{"2000-12-31", true},
		{"2000-02-29", true},
		{"2004-02-29", true},
		{"2020-04-30", true},
		// 5 datas diferentes inválidas
		{"200-01-01", false},
		{"2000-00-00", false},
		{"2000-11-31", false},
		{"2001-02-29", false},
		{"2023-02-29", false},
	}
	for _, c := range casos {
		if DataEn(c.data) != c.validade {
			t.Errorf("DataEn %s deveria ser %v", c.data, c.validade)
		}
	}
}

package main

import "testing"

func TestSoma(t *testing.T) {

	total := Soma(15, 10)

	if total != 25 {
		t.Errorf("Resultado errado: Resultado %d. Esperado: %d", total, 25)
	}
}

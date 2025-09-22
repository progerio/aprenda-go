package ola

import "testing"

func TestOla(t *testing.T) {
	obtido := Ola("Chris", "br")
	esperado := "Olá, Chris"

	if obtido != esperado {
		t.Errorf("obtido '%s' esperado '%s'", obtido, esperado)
	}
}

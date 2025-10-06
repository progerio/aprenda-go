package sync

import "testing"

func TestContador(t *testing.T) {
	t.Run("incrementar o contador 3 vezes no valor 3", func(t *testing.T) {
		contador := Contador{}
		contador.Incrementar()
		contador.Incrementar()
		contador.Incrementar()

		verificaContador(t, contador, 3)
	})
}

func verificaContador(t *testing.T, contador Contador, esperado int) {
	t.Helper()
	if contador.Valor() != esperado {
		t.Errorf("resultado %d, esperado %d", contador.Valor(), esperado)
	}
}

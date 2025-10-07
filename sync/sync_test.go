package sync

import (
	"sync"
	"testing"
)

func TestContador(t *testing.T) {
	t.Run("incrementar o contador 3 vezes no valor 3", func(t *testing.T) {
		contador := Contador{}
		contador.Incrementar()
		contador.Incrementar()
		contador.Incrementar()

		verificaContador(t, &contador, 3)
	})

	t.Run("Roda concorrentemente em segurança", func(t *testing.T) {
		contagemEsperada := 1000
		contador := NovoContador()

		var wg sync.WaitGroup
		wg.Add(contagemEsperada)

		for range contagemEsperada {
			go func(w *sync.WaitGroup) {
				contador.Incrementar()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		verificaContador(t, contador, contagemEsperada)
	})

}

func verificaContador(t *testing.T, contador *Contador, esperado int) {
	t.Helper()
	if contador.Valor() != esperado {
		t.Errorf("resultado %d, esperado %d", contador.Valor(), esperado)
	}
}

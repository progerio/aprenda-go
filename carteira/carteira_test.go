package carteira

import (
	"testing"
)

func TestCarteira(t *testing.T) {

	confirmaSaldo := func(t *testing.T, carteira Carteira, valorEsperado Bitcoin) {
		t.Helper()
		resultado := carteira.Saldo()

		if resultado != valorEsperado {
			t.Errorf("resultado %s, esperado %s", resultado, valorEsperado)
		}
	}

	confirmaErro := func(t *testing.T, resultado error, esperado error) {
		t.Helper()
		if resultado == nil {
			t.Fatal("esperava um erro mas não recebeu nenhum")
		}

		if resultado != esperado {
			t.Errorf("resultado de erro '%s', esperado '%s'", resultado, esperado)
		}

	}

	t.Run("Depositar", func(t *testing.T) {
		carteira := Carteira{}
		carteira.Depositar(Bitcoin(10))
		confirmaSaldo(t, carteira, Bitcoin(10))

	})
	t.Run("Retirar", func(t *testing.T) {
		carteira := Carteira{saldo: Bitcoin(20)}
		carteira.Retirar(Bitcoin(10))
		confirmaSaldo(t, carteira, Bitcoin(10))
	})

	t.Run("Retirar com saldo insuficiente", func(t *testing.T) {
		saldoInicial := Bitcoin(20)
		carteira := Carteira{saldo: saldoInicial}
		erro := carteira.Retirar(Bitcoin(100))

		confirmaSaldo(t, carteira, saldoInicial)
		confirmaErro(t, erro, ErroSaldoInsuficiente)
	})
}

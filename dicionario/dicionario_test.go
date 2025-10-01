package dicionario

import "testing"

func TestBusca(t *testing.T) {
	dicionario := Dicionario{"teste": "isso é apenas um teste"}

	t.Run("palavra existente", func(t *testing.T) {
		resultado, _ := dicionario.Busca("teste")
		esperado := "isso é apenas um teste"
		comparaStrings(t, resultado, esperado)
	})

	t.Run("palavra inexistente", func(t *testing.T) {
		_, err := dicionario.Busca("desconhecido")
		comparaErro(t, err, ErrNaoEncontrado)
	})
}

func comparaStrings(t *testing.T, resultado, esperado string) {
	t.Helper()
	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s' , dado %s", resultado, esperado, "teste")
	}
}

func comparaErro(t *testing.T, resultado, esperado error) {
	t.Helper()
	if resultado != esperado {
		t.Errorf("resultado erro '%s', esperado '%s'", resultado, esperado)
	}
}

func TestAdiciona(t *testing.T) {
	dicionario := Dicionario{}
	dicionario.Adiciona("teste", "isso é apenas um teste")

	esperado := "isso é apenas um teste"

	resultado, err := dicionario.Busca("teste")
	if err != nil {
		t.Fatal("não foi possível encontrar a palavra adicionada:", err)
	}

	if esperado != resultado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}

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
	t.Run("palavra nova", func(t *testing.T) {
		dicionario := Dicionario{}
		palavra := "teste"
		definicao := "isso é apenas um teste"
		err := dicionario.Adiciona(palavra, definicao)
		comparaErro(t, err, nil)
		comparaDefinicao(t, dicionario, palavra, definicao)
	})
	t.Run("palavra existente", func(t *testing.T) {
		palavra := "teste"
		definicao := "isso é apenas um teste"
		dicionario := Dicionario{palavra: definicao}
		err := dicionario.Adiciona(palavra, "teste novo")
		comparaErro(t, err, ErrPalavraExistente)
		comparaDefinicao(t, dicionario, palavra, definicao)
	})
}

func comparaDefinicao(t *testing.T, dicionario Dicionario, palavra, definicao string) {
	t.Helper()
	resultado, err := dicionario.Busca(palavra)
	if err != nil {
		t.Fatal("deveria ter encontrado palavra adicionada:", err)
	}
	if resultado != definicao {
		t.Errorf("resultado '%s', esperado '%s'", resultado, definicao)
	}
}
func TestUpdate(t *testing.T) {
	palavra := "teste"
	definicao := "isso é apenas um teste"
	dicionario := Dicionario{palavra: definicao}
	novaDefinicao := "nova definição"
	dicionario.Atualiza(palavra, novaDefinicao)
	comparaDefinicao(t, dicionario, palavra, novaDefinicao)
}

func TestDeleta(t *testing.T) {
	palavra := "teste"
	definicao := "isso é apenas um teste"
	dicionario := Dicionario{palavra: definicao}
	dicionario.Deleta(palavra)
	_, err := dicionario.Busca(palavra)
	if err != ErrNaoEncontrado {
		t.Errorf("espera-se que '%s' seja deletado", palavra)
	}
}

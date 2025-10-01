package dicionario

import "errors"

func Busca(dicionario map[string]string, palavra string) string {
	return dicionario[palavra]
}

var ErrNaoEncontrado = errors.New("não foi possível encontrar a palavra que você procura")

type Dicionario map[string]string

func (d Dicionario) Busca(palavra string) (string, error) {
	definicao, existe := d[palavra]
	if !existe {
		return "", ErrNaoEncontrado
	}
	return definicao, nil
}
func (d Dicionario) Adiciona(palavra, definicao string) {
	d[palavra] = definicao
}

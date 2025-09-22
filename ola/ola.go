package ola

import "fmt"

func Ola(nome, linguagem string) string {
	return fmt.Sprintf(
		"%s, %s",
		cumprimento(linguagem),
		nome,
	)
}

var comprimentos = map[string]string{
	"br": "Olá",
	"fr": "Bonjour",
}

func cumprimento(linguagem string) string {
	comprimento, existe := comprimentos[linguagem]

	if existe {
		return comprimento
	}
	return "Hello, "
}

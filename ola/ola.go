package ola

const prefixoOlaPortugues = "Olá, "

func Ola(nome string) string {
	return prefixoOlaPortugues + nome
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

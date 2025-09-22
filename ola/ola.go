package ola

func Ola(nome, linguagem string) string {
	if linguagem == "br" {
		return "Olá, " + nome
	}

	if linguagem == "fr" {
		return "Bonjour, " + nome
	}

	return "Hello, " + nome
}


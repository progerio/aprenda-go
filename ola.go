package main

import "fmt"

const espanhol = "espanhol"
const frances = "frances"

const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "

const prefixoOlaFrances = "Bonjour, "

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "mundo"
	}

	prefixo := prefixoOlaPortugues

	switch idioma {
	case frances:
		prefixo = prefixoOlaFrances
	case espanhol:
		prefixo = prefixoOlaEspanhol
	}

	return prefixo + nome

}

func main() {
	fmt.Println(Ola("mundo", "portugues"))
}

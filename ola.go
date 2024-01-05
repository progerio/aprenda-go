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

	if idioma == espanhol {
		return prefixoOlaEspanhol + nome
	}

	if idioma == frances {
		return prefixoOlaFrances + nome
	}
	return prefixoOlaPortugues + nome

}

func main() {
	fmt.Println(Ola("mundo", "portugues"))
}

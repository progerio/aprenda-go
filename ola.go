package main

import "fmt"

const espanhol = "espanhol"

const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "mundo"
	}

	if idioma == espanhol {
		return prefixoOlaEspanhol + nome
	}
	return prefixoOlaPortugues + nome

}

func main() {
	fmt.Println(Ola("mundo", "portugues"))
}

package main

import "fmt"

const prefixoOlaPortugues = "Olá, "

func Ola(nome string) string {
	if len(nome) > 0 {
		return prefixoOlaPortugues + nome
	}
	return prefixoOlaPortugues + "mundo"

}

func main() {
	fmt.Println(Ola("mundo"))
}

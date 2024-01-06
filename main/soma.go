package main

func Soma(numeros []int) int {
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}

	return soma
}

func SomaTudo(numerosParaSomar ...[]int) []int {
	quantidadeDeNumeros := len(numerosParaSomar)
	var somas []int

	somas = make([]int, quantidadeDeNumeros)
	for _, numeros := range numerosParaSomar {
		somas = append(somas, Soma(numeros))
	}

	return somas
}

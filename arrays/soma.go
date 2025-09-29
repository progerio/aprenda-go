package arrays

func Soma(numeros []int) int {
	resultado := 0
	for _, numero := range numeros {
		resultado += numero
	}
	return resultado
}

package iteracao

const quantidadeRepeticoes = 5

func Repetir(caracter string) string {
	var repeticoes string
	for i := 0; i < quantidadeRepeticoes; i++ {
		repeticoes += caracter
	}

	return repeticoes
}

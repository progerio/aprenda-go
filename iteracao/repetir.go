package iteracao

func Repetir(caracter string) string {
	var repeticoes string
	for i := 0; i < 5; i++ {
		repeticoes = repeticoes + caracter
	}

	return repeticoes
}

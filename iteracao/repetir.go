package iteracao

func Repetir(caractere string, quantidade int) string {
	var repeticoes string
	for range quantidade {
		repeticoes += caractere
	}
	return repeticoes
}

package concorrencia

type VerificadorWebSite func(string) bool
type resultado struct {
	string
	bool
}

func VerificadorWebSites(vw VerificadorWebSite, urls []string) map[string]bool {
	resultados := make(map[string]bool)
	canalResultado := make(chan resultado)

	for _, url := range urls {
		go func(u string) {
			canalResultado <- resultado{u, vw(u)}
		}(url)
	}

	for range urls {
		resultado := <-canalResultado
		resultados[resultado.string] = resultado.bool
	}
	return resultados
}

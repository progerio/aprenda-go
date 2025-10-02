package corredor

import (
	"net/http"
	"time"
)

func Corredor(a, b string) (vencedor string) {
	duracaoA := medirTempoDeResposta(a)
	duracaoB := medirTempoDeResposta(b)

	if duracaoA < duracaoB {
		return a
	}
	return b
}

func medirTempoDeResposta(url string) time.Duration {
	inicio := time.Now()
	http.Get(url)
	return time.Since(inicio)
}

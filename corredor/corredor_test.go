package corredor

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCorredor(t *testing.T) {
	t.Run("compara a velocidade de servidores, retornando o endereço do mais rápido", func(t *testing.T) {

		servidorLento := criarServidorcomAtraso(20 * time.Millisecond)
		servidorRapido := criarServidorcomAtraso(0 * time.Millisecond)

		servidorLento.Close()
		servidorRapido.Close()

		_, err := Corredor(servidorLento.URL, servidorRapido.URL)

		if err != nil {
			t.Error("esperava um erro, mas não obtive um")
		}
	})

	t.Run("retorna um erro se o tempo limite for excedido", func(t *testing.T) {
		servidor := criarServidorcomAtraso(25 * time.Millisecond)
		defer servidor.Close()

		_, err := Configuravel(servidor.URL, servidor.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("esperava um erro, mas não obtive um")
		}
	})
}

func criarServidorcomAtraso(duracao time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duracao)
		w.WriteHeader(http.StatusOK)
	}))
}

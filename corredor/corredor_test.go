package corredor

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCorredor(t *testing.T) {
	servidorLento := criarServidorcomAtraso(20 * time.Millisecond)
	servidorRapido := criarServidorcomAtraso(0 * time.Millisecond)

	servidorLento.Close()
	servidorRapido.Close()

	URLLenta := servidorLento.URL
	URLRapida := servidorRapido.URL

	esperado := URLRapida
	resultado := Corredor(URLLenta, URLRapida)
	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}

func criarServidorcomAtraso(duracao time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duracao)
		w.WriteHeader(http.StatusOK)
	}))
}

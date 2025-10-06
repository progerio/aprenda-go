package contexto

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type StubStore struct {
	response  string
	cancelled bool
}

func (s *StubStore) Fetch() string {
	return s.response
}
func (s *StubStore) Cancel() {
	s.cancelled = true
}

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Microsecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}
func TestHandler(t *testing.T) {
	data := "olá, mundo"
	svr := Server(&StubStore{response: data})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	svr.ServeHTTP(response, request)

	if response.Body.String() != data {
		t.Errorf("resultado %s , esperado %s ", response.Body.String(), data)
	}

	t.Run("avisa a store para cancelar o trabalho se a requisição for cancelada  ", func(t *testing.T) {
		store := &SpyStore{response: data}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)
		response := httptest.NewRecorder()
		svr.ServeHTTP(response, request)
		if !store.cancelled {
			t.Errorf("store não foi avisada para cancelar")
		}
	})

	t.Run("retorna dados da store", func(t *testing.T) {
		store := SpyStore{response: data}
		svr := Server(&store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
		svr.ServeHTTP(response, request)
		if response.Body.String() != data {
			t.Errorf(`resultado "%s", esperado "%s"`, response.Body.String(), data)
		}
		if store.cancelled {
			t.Error("não deveria ter cancelado a store")
		}
	})
}

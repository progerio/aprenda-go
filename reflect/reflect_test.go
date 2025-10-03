package reflect

import (
	"reflect"
	"testing"
)

func TestPercorre(t *testing.T) {
	esperado := "Chris"
	var resultado []string

	x := struct {
		Nome string
	}{esperado}

	percorre(x, func(entrada string) {
		resultado = append(resultado, entrada)
	})
	if resultado[0] != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado[0], esperado)
	}
}

func percorre(x interface{}, fn func(entrada string)) {
	valor := reflect.ValueOf(x)
	campo := valor.Field(0)
	fn(campo.String())
}

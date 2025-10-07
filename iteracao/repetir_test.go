package iteracao

import (
	"fmt"
	"testing"
)

func TestRepetir(t *testing.T) {
	repeticoes := Repetir("a", 5)
	esperado := "aaaaa"
	if repeticoes != esperado {
		t.Errorf("esperado '%s', mas obteve '%s", esperado, repeticoes)
	}
}

func BenchmarkRepetir(b *testing.B) {
	for b.Loop() {
		Repetir("a", 5)
	}
}

func ExampleRepetir() {
	repeticoes := Repetir("a", 5)
	fmt.Println(repeticoes)
	// Output: aaaaa
}

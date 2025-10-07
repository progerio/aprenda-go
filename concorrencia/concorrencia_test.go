package concorrencia

import (
	"reflect"
	"testing"
	"time"
)

func mockVerificadorWebSite(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func TestVerificadorWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}
	esperado := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	resultado := VerificadorWebSites(mockVerificadorWebSite, websites)

	if !reflect.DeepEqual(esperado, resultado) {
		t.Fatalf("esperado %v, recebido %v", esperado, resultado)
	}
}

func slowStubVerifiadorWebsite(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkVerificadorWebsites(b *testing.B) {
	urls := make([]string, 100)

	for i := range 100 {
		urls[i] = "http://www.google.com"
	}

	for i := 0; i < b.N; i++ {
		VerificadorWebSites(slowStubVerifiadorWebsite, urls)
	}
}

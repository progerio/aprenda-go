package server

import (
	"fmt"
	"net/http"
)

type Handler interface {
	ServerHttp(http.ResponseWriter, *http.Request)
}

func PlayerStore(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	fmt.Fprint(w, GetPlayerScore(player))
}

func GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	}
	if name == "Floyd" {
		return "10"
	}
	return ""
}

package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func GameIndex(w http.ResponseWriter, _ *http.Request) {
	games := Games{
		Game{Name: "Call of Duty: Modern Warfare"},
		Game{Name: "Street Fighter"},
	}

	json.NewEncoder(w).Encode(games)
}

func GameShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameID := vars["gameId"]
	fmt.Fprintln(w, "Game show:", gameID)
}

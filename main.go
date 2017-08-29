package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/games", gameIndex)
	router.HandleFunc("/games/{gameId}", gameShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func gameIndex(w http.ResponseWriter, _ *http.Request) {
	games := games{
		game{Name: "Call of Duty: Modern Warfare"},
		game{Name: "Street Fighter"},
	}

	json.NewEncoder(w).Encode(games)
}

func gameShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameID := vars["gameId"]
	fmt.Fprintln(w, "Game show:", gameID)
}

type game struct {
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
	Rating    int    `json:"rating"`
}

type games []game

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/games", GameIndex)
	router.HandleFunc("/games/{gameId}", GameShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func GameIndex(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Game Index!")
}

func GameShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameId := vars["gameId"]
	fmt.Fprintln(w, "Game show:", gameId)
}

package main

type Game struct {
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
	Rating    int    `json:"rating"`
}

type Games []Game
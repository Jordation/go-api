package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type NewInput struct {
	Input string
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Handled.")
	json.NewEncoder(w).Encode("Hello World")
}

func HandleInput(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var input NewInput
	_ = json.NewDecoder(r.Body).Decode(&input)
	fmt.Println(input)

	if input.Input == "foo" {
		json.NewEncoder(w).Encode("bar")
	} else {
		json.NewEncoder(w).Encode("baz")
	}
}

func StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/", DefaultHandler)
	r.HandleFunc("/anything", HandleInput)
	fmt.Println("Server started on port 9000")
	log.Fatal(http.ListenAndServe(":9000", r))
}

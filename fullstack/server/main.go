package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aldarisbm/bachataistakingover/structs"
	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "Test Test Test")
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	p := structs.Health{
		Message:    "OK",
		StatusCode: 200,
	}

	json.NewEncoder(w).Encode(p)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/healthz", health).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aldarisbm/bachataistakingover/structs"
	"github.com/biter777/countries"
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

	must(json.NewEncoder(w).Encode(p))
}

func countriesList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	countryList := countries.AllInfo()
	must(json.NewEncoder(w).Encode(countryList))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/healthz", health).Methods("GET")
	r.HandleFunc("/api/v1/countries", countriesList).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", r))
}

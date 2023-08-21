package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/", helloHandler).Methods("GET")
	r.HandleFunc("/greet/{name}", greetHandler).Methods("GET")

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(
		":%s",
		os.Getenv("SERVER_PORT"),
	), r))
}


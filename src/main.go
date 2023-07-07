package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", helloHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(
		":%s",
		os.Getenv("SERVER_PORT"),
	), nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, New World!"))
}

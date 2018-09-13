package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	port := "8080"

	http.HandleFunc("/", StatusHandler)

	log.Printf("starting server on %s\n", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Status: OK")
}

package main

import (
	"errors"
	"io"
	"log"
	"net/http"
)

func handleStatus(w http.ResponseWriter, r *http.Request) {
	log.Print("handling status request")
	io.WriteString(w, "Server is running\n")
}

func main() {
	http.HandleFunc("/status", handleStatus)

	err := http.ListenAndServe(":3333", nil)
	if err != nil && errors.Is(err, http.ErrServerClosed) {
		log.Print("closing server")
	} else {
		log.Fatal(err)
	}
}

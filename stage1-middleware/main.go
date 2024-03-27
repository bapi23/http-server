package main

import (
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

func logTimeElapsed(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		f(w, r)
		end := time.Now()

		log.Printf("Handling: %s, took: %s", r.URL.Path, end.Sub(start))
	}
}

func handleStatus(w http.ResponseWriter, r *http.Request) {
	log.Print("handling status request")
	io.WriteString(w, "Server is running\n")
}

func main() {
	http.HandleFunc("/status", logTimeElapsed(handleStatus))

	err := http.ListenAndServe(":3333", nil)
	if err != nil && errors.Is(err, http.ErrServerClosed) {
		log.Print("closing server")
	} else {
		log.Fatal(err)
	}
}

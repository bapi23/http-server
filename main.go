package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("foobar"))
	w.WriteHeader(http.StatusBadRequest)
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)

	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	server := http.Server{
		Addr:    ":3333",
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		log.Print("shutting down server")
		shoutDownCtx, shutdownRelease := context.WithTimeout(context.Background(), time.Second*10)
		defer shutdownRelease()
		server.Shutdown(shoutDownCtx)
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		err := server.ListenAndServe()
		if err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Print("closing server")
		} else {
			log.Fatal(err)
		}
		wg.Done()
	}()

	wg.Wait()
}

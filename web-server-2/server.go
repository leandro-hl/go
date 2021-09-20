package main

import (
	"context"
	"hl/web/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	os.Setenv("APP_PORT", ":3002")

	l := log.New(os.Stdout, "api:", log.LstdFlags)

	gh := handlers.NewBye(l)
	hh := handlers.NewHello(l)
	ph := handlers.NewProductsController(l)

	// create a new server mux and register the handlers
	// Mutex is a mutual exclusion lock
	// Here, Mux is multiplexer
	sm := http.NewServeMux()
	sm.Handle("/bye", gh)
	sm.Handle("/hello", hh)
	sm.Handle("/products", ph)

	s := &http.Server{
		Addr:         os.Getenv("APP_PORT"),
		Handler:      sm,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second, // Max time for connections using TCP keep-alive.
	}

	// Start server
	go func() {
		l.Println("Starting server...")

		err := s.ListenAndServe()

		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interrupt and gracefully shutdown the server
	// buffered channel with capacity for 1 event
	c := make(chan os.Signal, 1)

	// Channel listens for Interrupt or Kill events
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal: ", sig)

	// Gracefully shutdown the server, waiting max 30 seconds for the current operations to complete.
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}

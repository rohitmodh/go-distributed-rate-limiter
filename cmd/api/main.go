package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mux := http.NewServeMux()

	health := chain(http.HandlerFunc(healthHandler), wrap)
	home := chain(http.HandlerFunc(homeHandler), wrap)

	mux.Handle("/health", health)
	mux.Handle("/", home)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	signalCtx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	if err := runServer(server, signalCtx); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

func runServer(server *http.Server, signalCtx context.Context) error {
	//
	log.Println("Starting server on : 8080")

	serverErrors := make(chan error, 1)

	go func() {
		serverErrors <- server.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
	case <-signalCtx.Done():
		log.Println("Shutdown signal received")

		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			5*time.Second,
		)
		defer cancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Printf("server shutdown error: %v", err)
			return err
		}
	}
	return nil
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside home")
	w.Header().Set("content-type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"service": "go-backend-platform",
		"version": "v1",
	})
}

func wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		defer func() {
			fmt.Println("method = ", r.Method, " path = ", r.RequestURI, "completed in", time.Since(start))
		}()
		next.ServeHTTP(w, r)
	})
}

func chain(handler http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

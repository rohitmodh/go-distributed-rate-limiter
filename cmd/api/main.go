package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	health := chain(http.HandlerFunc(healthHandler), wrap)
	home := chain(http.HandlerFunc(homeHandler), wrap)

	mux.Handle("/health", health)
	mux.Handle("/", home)

	log.Println("Starting server on : 8080")

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal(err)
	}
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

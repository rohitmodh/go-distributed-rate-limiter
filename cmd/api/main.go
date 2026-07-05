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

	mux.HandleFunc("/health", wrap(healthHandler))
	mux.HandleFunc("/", wrap(homeHandler))
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

func wrap(fn func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		defer func() {
			fmt.Println("method = ", r.Method, " path = ", r.RequestURI, "completed in", time.Since(start))
		}()
		fn(w, r)
	}
}

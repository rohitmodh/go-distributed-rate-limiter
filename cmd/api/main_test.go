package main

import (
	"context"
	"net/http"
	"testing"
	"time"
)

func TestRunServerShutdown(t *testing.T) {
	signalCtx, cancel := context.WithCancel(
		context.Background(),
	)
	defer cancel()

	server := &http.Server{
		Addr: ":0",
	}

	serverErr := make(chan error, 1)

	go func() {
		serverErr <- runServer(server, signalCtx)
	}()

	cancel()

	select {
	case err := <-serverErr:
		if err != nil {
			t.Fatalf("expected clean shutdown, got error: %v", err)
		}

	case <-time.After(2 * time.Second):
		t.Fatal("expected server to shut down gracefully, but it didn't")
	}
}

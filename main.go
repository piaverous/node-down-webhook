package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/piaverous/node-down-webhook/handlers"
)

var (
	httpAddr = flag.String("http", ":8080", "Listen address")
)

func run() error {
	flag.Parse()

	http.HandleFunc("/healthz", handlers.Healthz)
	http.HandleFunc("/webhook", handlers.Webhook)

	log.Printf("Listening on %s...", *httpAddr)
	if err := http.ListenAndServe(*httpAddr, nil); err != nil {
		return fmt.Errorf("unable to listen for requests: %w", err)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

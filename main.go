package main

import (
	_ "embed"
	"log"
	"net/http"

	"github.com/gdwr/chaoss/internal/middleware"
)

//go:embed docs/openapi.yml
var openapi string

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /openapi.yml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/yml")
		w.Write([]byte(openapi))
	})
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	wrappedMux := middleware.NewLogger(mux)
	log.Fatal(http.ListenAndServe(":8080", wrappedMux))
}

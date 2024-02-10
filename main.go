package main

import (
	_ "embed"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gdwr/chaoss/internal/middleware"
	"github.com/gdwr/chaoss/internal/repository"
	"github.com/gdwr/chaoss/internal/schemas"
)

//go:embed docs/openapi.yml
var openapi string

type GetMatchResponse struct {
	*schemas.Match
	Token schemas.Guid `json:"token"`
}

func main() {
	matchRepository := repository.NewInMemoryMatchRepository()
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

	mux.HandleFunc("GET /match", func(w http.ResponseWriter, r *http.Request) {
		match, err := matchRepository.RandomMatch()
		if err != nil {
			match = matchRepository.NewMatch()
		}

		response := GetMatchResponse{
			Match: match,
			Token: schemas.NewGuid(),
		}

		content, err := json.Marshal(&response)
		if err != nil {
			log.Printf("%s", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(content)
	})
	wrappedMux := middleware.NewLogger(mux)
	log.Printf("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", wrappedMux))
}

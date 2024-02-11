package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gdwr/chaoss/internal/middleware"
	"github.com/gdwr/chaoss/internal/repository"
	"github.com/gdwr/chaoss/internal/schemas"
)

type GetMatchResponse struct {
	*schemas.Match
	Token schemas.Guid `json:"token"`
}

func main() {
	logger := log.New(log.Writer(), "chaoss: ", log.Ltime|log.LUTC|log.Lmsgprefix)
	matchRepository := repository.NewInMemoryMatchRepository()
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "website/index.html")
	})
	mux.HandleFunc("GET /openapi.yml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/yml")
		http.ServeFile(w, r, "docs/openapi.yml")
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
	mux.HandleFunc("GET /match/{guid}", func(w http.ResponseWriter, r *http.Request) {
		guid := r.PathValue("guid")
		match, err := matchRepository.GetMatch(guid)
		if err != nil {
			http.Error(w, "Match not found", http.StatusNotFound)
			return
		}

		content, err := json.Marshal(match)
		if err != nil {
			log.Printf("%s", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(content)
	})

	wrappedMux := middleware.NewLogger(logger, mux)
	logger.Printf("Listening on :8080")
	logger.Fatal(http.ListenAndServe(":8080", wrappedMux))
}

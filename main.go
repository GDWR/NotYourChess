package main

import (
	"embed"
	"encoding/json"
	"log"
	"net/http"
	"path"

	"github.com/gdwr/chaoss/internal/middleware"
	"github.com/gdwr/chaoss/internal/repository"
	"github.com/gdwr/chaoss/internal/schemas"
)

type GetMatchResponse struct {
	*schemas.Match
	Token schemas.Guid `json:"token"`
}

//go:embed website/dist/*
var website embed.FS

func main() {
	logger := log.New(log.Writer(), "chaoss: ", log.Ltime|log.LUTC|log.Lmsgprefix)
	matchRepository := repository.NewInMemoryMatchRepository()
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		// Serve content from embed website
		file := path.Clean(r.URL.Path)
		if file == "/" {
			file = "/index.html"
		}

		content, err := website.ReadFile("website/dist" + file)
		if err != nil {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		// Determine content type by file extension
		contentType := "text/plain"
		switch path.Ext(file) {
		case ".html":
			contentType = "text/html"
		case ".css":
			contentType = "text/css"
		case ".js":
			contentType = "application/javascript"
		case ".svg":
			contentType = "image/svg+xml"
		}

		w.Header().Set("Content-Type", contentType)
		w.Write(content)
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
	mux.HandleFunc("GET /api/match", func(w http.ResponseWriter, r *http.Request) {
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

		http.SetCookie(w, &http.Cookie{
			Name:  "matchSession",
			Value: response.Token,
		})
		w.Header().Set("Content-Type", "application/json")
		w.Write(content)
	})
	mux.HandleFunc("GET /api/match/{guid}", func(w http.ResponseWriter, r *http.Request) {
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

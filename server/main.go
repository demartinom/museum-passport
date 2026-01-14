package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/demartinom/museum-passport/cache"
	"github.com/demartinom/museum-passport/handlers"
	"github.com/demartinom/museum-passport/museums"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Initialize cache
	cache := cache.NewCache()

	// Creates map of museum clients to be used by different handlers
	clients := map[string]museums.Client{
		"met":     museums.NewMetClient(cache),
		"harvard": museums.NewHarvardClient(os.Getenv("HARVARD_KEY"), cache),
	}
	ArtworkHandler := handlers.NewArtworkHandler(clients)
	SearchHandler := handlers.NewSearchHandler(clients)

	r := chi.NewRouter()
	r.Get("/api/artwork/{id}", ArtworkHandler.GetArtwork)
	r.Get("/api/search", SearchHandler.SearchArtwork)

	port := os.Getenv("PORT")
	fmt.Printf("Starting server at port%s\n", port)
	http.ListenAndServe(port, r)
}

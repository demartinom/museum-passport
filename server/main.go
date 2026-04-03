package main

import (
	"log"
	"net/http"
	"os"

	"github.com/demartinom/museum-passport/ai"
	"github.com/demartinom/museum-passport/cache"
	"github.com/demartinom/museum-passport/handlers"
	"github.com/demartinom/museum-passport/museums"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("DOCKER_ENV") == "" {
		if err := godotenv.Load(); err != nil {
			// Only log this if we aren't in a container
			log.Println("Running in local mode without .env file")
		}
	}

	// Initialize cache
	cache := cache.NewCache()

	// Create museum clients
	clients := map[string]museums.Client{
		"met":     museums.NewMetClient(cache),
		"harvard": museums.NewHarvardClient(os.Getenv("HARVARD_KEY"), cache),
	}

	ArtworkHandler := handlers.NewArtworkHandler(clients)
	SearchHandler := handlers.NewSearchHandler(clients)

	// Create AI client
	openAIKey := os.Getenv("OPENAI_KEY")
	if openAIKey == "" {
		log.Fatal("OPENAI_KEY is not set")
	}
	summaryClient := ai.NewSummaryClient(openAIKey)
	summaryHandler := handlers.NewSummaryHandler(summaryClient, cache, clients)

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "https://museum-passport.vercel.app"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/api/artwork/{id}", ArtworkHandler.GetArtwork)
	r.Get("/api/search", SearchHandler.SearchArtwork)
	r.Get("/api/summary", summaryHandler.GenerateSummary)

	// Fly assigns a dynamic port
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001" // local fallback
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

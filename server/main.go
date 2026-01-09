package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// metClient := museums.NewMetClient()
	// harvardClient := museums.NewHarvardClient(os.Getenv("HARVARD_KEY"))

	r := chi.NewRouter()

	port := os.Getenv("PORT")
	fmt.Printf("Starting server at port%s\n", port)
	http.ListenAndServe(port, r)
}

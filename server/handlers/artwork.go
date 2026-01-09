package handlers

import "github.com/demartinom/museum-passport/museums"

type ArtworkHandler struct {
	Clients map[string]museums.Client
}

func NewArtworkHandler(clients map[string]museums.Client) *ArtworkHandler {
	return &ArtworkHandler{Clients: clients}
}

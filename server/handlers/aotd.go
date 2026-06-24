package handlers

import (
	"net/http"

	"github.com/demartinom/museum-passport/cache"
)

type AOTDHandler struct {
	Cache *cache.Cache
}

func NewAOTDHandler(c *cache.Cache) *AOTDHandler {
	return &AOTDHandler{Cache: c}
}


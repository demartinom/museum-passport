package museums

import "github.com/demartinom/museum-passport/cache"

type PrincetonClient struct {
	BaseURL string
	Cache   *cache.Cache
}

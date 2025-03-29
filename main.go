package main

import (
	"github.com/mattr/pokedex/internal/api"
	"github.com/mattr/pokedex/internal/cache"
	"time"
)

func main() {
	cfg := &config{Cache: cache.NewCache(30 * time.Second), Pokedex: make(map[string]api.Pokemon)}
	startRepl(cfg)
}

package config

import (
	"fmt"
	"github.com/nyumbapoa/backend/internal/store"
)

type Config struct {
}

func (cfg *Config) GetStore() (store.Store, error) {

	return nil, fmt.Errorf("config: unknown store type: %s", cfg)
}

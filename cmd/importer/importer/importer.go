package importer

import (
	"github.com/superhero-choice-importer/internal/cache"
	"github.com/superhero-choice-importer/internal/config"
	"github.com/superhero-choice-importer/internal/db"
)

// Importer holds all the data relevant.
type Importer struct {
	DB    *db.DB
	Cache *cache.Cache
}

// NewImporter configures Importer.
func NewImporter(cfg *config.Config) (im *Importer, err error) {
	dbs, err := db.NewDB(cfg)
	if err != nil {
		return nil, err
	}

	ch, err := cache.NewCache(cfg)
	if err != nil {
		return nil, err
	}

	return &Importer{
		DB:    dbs,
		Cache: ch,
	}, nil
}

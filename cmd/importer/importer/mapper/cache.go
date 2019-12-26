package mapper

import (
	cache "github.com/superhero-choice-importer/internal/cache/model"
	"github.com/superhero-choice-importer/internal/db/model"
)

// MapDBChoicesToCache maps DB Choice models to Cache Choice models.
func MapDBChoicesToCache(chs []model.Choice) (choices []cache.Choice) {
	for _, c := range chs {
		choices = append(choices, cache.Choice{
			ID:                c.ID,
			Choice:            c.Choice,
			SuperheroID:       c.SuperheroID,
			ChosenSuperheroID: c.ChosenSuperheroID,
			CreatedAt:         c.CreatedAt,
		})
	}

	return choices
}

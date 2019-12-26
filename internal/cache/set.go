package cache

import (
	"fmt"
	"github.com/superhero-choice-importer/internal/cache/model"
	"time"
)

// SetChoices stores choice(like only, dislikes only go to DB) into Redis cache.
func (c *Cache) SetChoices(choices []model.Choice) error {
	for _, choice := range choices {
		err := c.Redis.Set(
			fmt.Sprintf("%s.%s", choice.SuperheroID, choice.ChosenSuperheroID),
			choice,
			time.Hour*time.Duration(24)*time.Duration(7),
		).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

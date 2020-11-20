/*
  Copyright (C) 2019 - 2021 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package cache

import (
	"fmt"
	"github.com/superhero-match/superhero-choice-importer/internal/cache/model"
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

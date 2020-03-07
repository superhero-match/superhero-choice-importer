/*
  Copyright (C) 2019 - 2020 MWSOFT
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
package importer

import (
	"github.com/superhero-choice-importer/cmd/importer/importer/mapper"
)

// Import import Superheros data from DB to Cache.
func (i *Importer) Import() error {
	offset := int64(0)

	for {
		// 1. Fetch the first batch of Choice from DB.
		choicesDB, err := i.DB.GetChoices(offset)
		if err != nil {
			return err
		}

		if len(choicesDB) == 0 {
			break
		}

		// 2. Set offset for the next batch.
		offset += int64(len(choicesDB))

		choicesCache := mapper.MapDBChoicesToCache(choicesDB)

		err = i.Cache.SetChoices(choicesCache)
		if err != nil {
			return err
		}

		// 3. Check if the len of batch returned is less than the batch size,
		// if so, it means it was the last iteration, after the Choice
		// have been saved to Cache the loop can be broken.
		if len(choicesDB) < i.DB.Limit {
			break
		}
	}

	return nil
}

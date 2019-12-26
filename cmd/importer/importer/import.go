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

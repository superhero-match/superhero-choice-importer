package db

import (
	"github.com/superhero-choice-importer/internal/db/model"
)

// GetChoices fetches a batch of choice.
func (db *DB) GetChoices(offset int64) (choices []model.Choice, err error) {
	err = db.stmtGetChoices.Select(&choices, offset, db.Limit)
	if err != nil {
		return nil, err
	}

	return choices, nil
}

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
package db

import (
	"github.com/superhero-match/superhero-choice-importer/internal/db/model"
)

// GetChoices fetches a batch of choice.
func (db *DB) GetChoices(offset int64) (choices []model.Choice, err error) {
	err = db.stmtGetChoices.Select(&choices, offset, db.Limit)
	if err != nil {
		return nil, err
	}

	return choices, nil
}

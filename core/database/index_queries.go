package database

import (
	"social-network/view/index"
)

// UsersByNickname Gibt einen Index zurück, der Benutzer anhand ihrer Nicknames sortiert.
func (db *Database) UsersByNickname() *index.Index {
	idx := index.ByNickname()

	for _, u := range db.Users {
		idx.Insert(u)
	}

	return idx
}

// UsersByName Gibt einen Index zurück, der Benutzer anhand ihrer Namen sortiert.
func (db *Database) UsersByName() *index.Index {
	idx := index.ByName()

	for _, u := range db.Users {
		idx.Insert(u)
	}

	return idx
}

// UsersBySurname Gibt einen Index zurück, der Benutzer anhand ihrer Nachnamen sortiert.
func (db *Database) UsersBySurname() *index.Index {
	idx := index.BySurname()

	for _, u := range db.Users {
		idx.Insert(u)
	}

	return idx
}

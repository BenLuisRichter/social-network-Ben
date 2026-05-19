package database

import (
	"social-network/core/user"
	"social-network/view/index"
)

// Database repräsentiert unsere Datenbank, die alle Benutzer enthält.
type Database struct {
	Users []*user.User
}

// New erstellt eine neue leere Datenbank.
func New() *Database {
	return &Database{
		Users: []*user.User{},
	}
}

// AddUser fügt einen neuen Benutzer zur Datenbank hinzu.
func (db *Database) AddUser(u *user.User) {
	db.Users = append(db.Users, u)
}

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

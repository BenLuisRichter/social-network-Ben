package database

import (
	"social-network/core/user"
)

// GetDirectContacts liefert eine Liste mit den direkten Kontakten eines Benutzers zurück.
// Parameter:
//   - `id`: Die ID des Benutzers, dessen direkte Kontakte abgerufen werden sollen.
func (db *Database) GetDirectContacts(id string) []user.User {
	result := []user.User{}

	users := db.UsersById()
	e := users.Find(id)
	if e == nil {
		return result
	}

	for _, contactID := range e.User.Contacts {
		if c := users.Find(contactID); c != nil {
			result = append(result, *c.User)
		}
	}

	return result
}

// GetDirectContactIds liefert eine Liste mit den IDs aller direkten Kontakte eines Benutzers zurück.
// Parameter:
//   - `id`: Die ID des Benutzers, dessen direkte Kontakt-IDs abgerufen werden sollen.
func (db *Database) GetDirectContactIds(id string) []string {
	contacts := db.GetDirectContacts(id)
	contact_ids := []string{}

	for _, c := range contacts {
		contact_ids = append(contact_ids, c.Id)
	}

	return contact_ids
}

// GetContacts liefert eine Liste mit allen direkten und indirekten Kontakten eines Benutzers zurück.
// Parameter:
//   - `id`: Die ID des Benutzers, dessen Kontakte abgerufen werden sollen.
//   - `depth`: Die Anzahl der Ebenen von Kontakten, die berücksichtigt werden sollen.
//     Eine Tiefe von 1 bedeutet nur direkte Kontakte, 2 bedeutet direkte Kontakte und deren Kontakte usw.
func (db *Database) GetContacts(id string, depth int) []user.User {
	result := []user.User{}

	if depth < 1 {
		return result
	}

	indirect_contacts := []user.User{}
	visited := make(map[string]bool)
	visited[id] = true

	for _, c := range db.GetDirectContacts(id) {
		visited[c.Id] = true
		result = append(result, c)

		indirect_contacts = append(indirect_contacts, db.GetContacts(c.Id, depth-1)...)
	}

	for _, c := range indirect_contacts {
		if !visited[c.Id] {
			visited[c.Id] = true
			result = append(result, c)
		}
	}

	return result
}

// GetContactIds liefert eine Liste mit den IDs aller direkten und indirekten Kontakte eines Benutzers zurück.
// Parameter:
//   - `id`: Die ID des Benutzers, dessen Kontakt-IDs abgerufen werden sollen.
//   - `depth`: Die Anzahl der Ebenen von Kontakten, die berücksichtigt werden sollen.
//     Eine Tiefe von 1 bedeutet nur direkte Kontakte, 2 bedeutet direkte Kontakte und deren Kontakte usw.
func (db *Database) GetContactIds(id string, depth int) []string {
	contacts := db.GetContacts(id, depth)
	contact_ids := []string{}

	for _, c := range contacts {
		contact_ids = append(contact_ids, c.Id)
	}

	return contact_ids
}

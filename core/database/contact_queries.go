package database

import (
	"social-network/core/user"
)

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

	users := db.UsersById()
	e := users.Find(id)
	if e == nil {
		return result
	}

	u := e.User()
	direct_contacts := []user.User{}
	indirect_contacts := []user.User{}
	visited := make(map[string]bool)
	visited[u.Id] = true

	for _, contactID := range u.Contacts {
		if c := users.Find(contactID); c != nil {
			direct_contacts = append(direct_contacts, *c.User())
		}
	}

	for _, c := range direct_contacts {
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

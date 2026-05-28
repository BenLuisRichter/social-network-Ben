package database

import (
	"social-network/core/user"
)

// GetDirectContacts liefert eine Liste mit den direkten Kontakten eines Benutzers zurück.
// Parameter:
//   - `id`: Die ID des Benutzers, dessen direkte Kontakte abgerufen werden sollen.
func (db *Database) GetDirectContacts(id string) []user.User {
	// Hinweis:
	// - Holen Sie sich von `db` einen Suchindex, der die Benutzer
	//   nach ihrer ID gruppiert enthält.
	// - Verwenden Sie dann die `Find`-Methode dieses Suchindex,
	//   um den `User` zu bestimmen.
	// - Durchlaufen Sie dann die Kontaktliste dieses `User`, suchen Sie in `db`
	//   nach den IDs und kopieren Sie diese in das Ergebnis.

	result := []user.User{}

	// TODO

	return result
}

// GetDirectContactIds liefert eine Liste mit den IDs aller direkten Kontakte eines Benutzers zurück.
// Parameter:
//   - `id`: Die ID des Benutzers, dessen direkte Kontakt-IDs abgerufen werden sollen.
func (db *Database) GetDirectContactIds(id string) []string {
	// Hinweis:
	// - Verwenden Sie die `GetDirectContacts`-Methode, um die direkten Kontakte zu erhalten.
	//   Extrahieren Sie dann die IDs in einer Schleife.

	contact_ids := []string{}

	// TODO

	return contact_ids
}

// GetContacts liefert eine Liste mit allen direkten und indirekten Kontakten eines Benutzers zurück.
// Parameter:
//   - `id`: Die ID des Benutzers, dessen Kontakte abgerufen werden sollen.
//   - `depth`: Die Anzahl der Ebenen von Kontakten, die berücksichtigt werden sollen.
//     Eine Tiefe von 1 bedeutet nur direkte Kontakte, 2 bedeutet direkte Kontakte und deren Kontakte usw.
func (db *Database) GetContacts(id string, depth int) []user.User {
	// Hinweis:
	// - Verwenden Sie die `GetDirectContacts`-Methode, um die direkten Kontakte zu erhalten.
	// - Erstellen und befüllen Sie eine weitere Liste mit den indirekten Kontakten,
	//   indem Sie die `GetContacts`-Methode rekursiv auf jedem direkten Kontakt aufrufen.
	// - Um Duplikate und Selbstbezüge zu vermeiden, ist es sinnvoll, sich zu merken,
	//   welche User-IDs bereits besucht wurden. Dafür kann z.B. eine `map[string]bool` verwendet werden.

	result := []user.User{}

	// TODO

	return result
}

// GetContactIds liefert eine Liste mit den IDs aller direkten und indirekten Kontakte eines Benutzers zurück.
// Parameter:
//   - `id`: Die ID des Benutzers, dessen Kontakt-IDs abgerufen werden sollen.
//   - `depth`: Die Anzahl der Ebenen von Kontakten, die berücksichtigt werden sollen.
//     Eine Tiefe von 1 bedeutet nur direkte Kontakte, 2 bedeutet direkte Kontakte und deren Kontakte usw.
func (db *Database) GetContactIds(id string, depth int) []string {
	// Hinweis:
	// - Verwenden Sie die `GetContacts`-Methode, um alle Kontakte zu erhalten.
	// - Extrahieren Sie dann die IDs in einer Schleife.

	contact_ids := []string{}

	// TODO

	return contact_ids
}

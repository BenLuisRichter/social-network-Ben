package index

import (
	"fmt"
	"social-network/core/user"
	"social-network/view/element"
)

// Insert fügt einen neuen Benutzer in den Index ein.
func (idx *Index) Insert(u *user.User) {
	// Hinweis:
	// - Kopieren Sie sich die Wurzel in einen Pointer, den Sie verändern können.
	// - Steigen Sie dann mit einer Schleife in den Baum hinab, bis Sie ein leeres Element erreichen.
	// - Setzen Sie dann den Benutzer in diesem Element mit `SetUser`.
	// - Beim Abstieg müssen Sie entscheiden, ob Sie nach links oder rechts gehen,
	//   abhängig von der Sortierreihenfolge des Index.
	//   Diese Sortierreihenfolge wird durch die `key`-Funktion des Index bestimmt,
	//   die Sie mit `idx.key(e)` aufrufen können.
	//   Genauer: Die `key`-Funktion bestimmt für einen `User` den Schlüssel, nach dem sortiert wird.

	current := idx.root

	for !current.IsEmpty() {
		key := idx.key(u)
		currentKey := idx.key(current.User)

		if key < currentKey {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	current.SetUser(u)
}

// Find sucht einen Benutzer im Index anhand eines Schlüssels.
func (idx *Index) Find(key string) *element.Element {
	// Hinweis:
	// - Kopieren Sie sich die Wurzel in einen Pointer, den Sie verändern können.
	// - Steigen Sie dann mit einer Schleife in den Baum hinab, bis Sie ein leeres Element erreichen oder den gesuchten Schlüssel finden.
	// - Beim Abstieg müssen Sie entscheiden, ob Sie nach links oder rechts gehen,
	//   abhängig von der Sortierreihenfolge des Index.
	//   Diese Sortierreihenfolge wird durch die `key`-Funktion des Index bestimmt,
	//   die Sie mit `idx.key(e)` aufrufen können.

	current := idx.root

	for !current.IsEmpty() {
		currentKey := idx.key(current.User)

		if key == currentKey {
			return current
		} else if key < currentKey {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	return nil
}

// List liefert alle Benutzer im Index in sortierter Reihenfolge.
func (idx *Index) List() []*user.User {
	// Hinweis:
	// - Rufen Sie die `List`-Methode der Wurzel auf, um alle Benutzer in sortierter Reihenfolge zu erhalten.

	return idx.root.List()
}

// Keys liefert eine Liste aller Schlüssel im Index in sortierter Reihenfolge.
func (idx *Index) Keys() []string {
	// Hinweis:
	// - Rufen Sie die `List`-Methode des Index auf, um alle Benutzer in sortierter Reihenfolge zu erhalten.
	// - Extrahieren Sie dann die Schlüssel mit der `key`-Funktion des Index in einer Schleife.

	keys := []string{}

	for _, u := range idx.List() {
		keys = append(keys, idx.key(u))
	}

	return keys
}

// String gibt eine menschenlesbare Listen-Darstellung aller Schlüssel im Index zurück.
func (idx *Index) String() string {
	// Hinweis:
	// - Rufen Sie die `Keys`-Methode des Index auf, um alle Schlüssel in sortierter Reihenfolge zu erhalten.
	// - Verwandeln Sie diese Liste mittels `fmt.Sprintf` in einen String.

	return fmt.Sprintf("%v", idx.Keys())
}

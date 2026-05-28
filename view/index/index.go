package index

import (
	"fmt"
	"social-network/core/user"
	"social-network/view/element"
)

// Index repräsentiert einen binären Suchbaum, der Benutzer enthält.
// Der Baum dient als Suchindex, um Benutzer schnell anhand eines Schlüssels zu finden.
type Index struct {
	root *element.Element
	key  func(e *user.User) string
}

// New erstellt einen neuen leeren Index mit der gegebenen Schlüsselfunktion.
func New(key func(e *user.User) string) *Index {
	return &Index{
		root: element.Empty(),
		key:  key,
	}
}

// ById erstellt einen Index, der Benutzer anhand ihrer ID sortiert.
func ById() *Index {
	return New(func(u *user.User) string {
		return u.Id
	})
}

// ByNickname erstellt einen Index, der Benutzer anhand ihres Nickname sortiert.
func ByNickname() *Index {
	return New(func(u *user.User) string {
		return u.Nickname
	})
}

// ByName erstellt einen Index, der Benutzer anhand ihres Namens sortiert.
func ByName() *Index {
	return New(func(u *user.User) string {
		return u.Name
	})
}

// BySurname erstellt einen Index, der Benutzer anhand ihres Nachnamens sortiert.
func BySurname() *Index {
	return New(func(u *user.User) string {
		return u.Surname
	})
}

// Insert fügt einen neuen Benutzer in den Index ein.
func (idx *Index) Insert(e *user.User) {
	current := idx.root

	for !current.IsEmpty() {
		key := idx.key(e)
		currentKey := idx.key(current.User())

		if key < currentKey {
			current = current.LeftChild()
		} else {
			current = current.RightChild()
		}
	}

	current.SetUser(e)
}

// Find sucht einen Benutzer im Index anhand eines Schlüssels.
func (idx *Index) Find(key string) *element.Element {
	current := idx.root

	for !current.IsEmpty() {
		currentKey := idx.key(current.User())

		if key == currentKey {
			return current
		} else if key < currentKey {
			current = current.LeftChild()
		} else {
			current = current.RightChild()
		}
	}

	return nil
}

// List liefert alle Benutzer im Index in sortierter Reihenfolge.
func (idx *Index) List() []*user.User {
	return idx.root.List()
}

// Keys liefert eine Liste aller Schlüssel im Index in sortierter Reihenfolge.
func (idx *Index) Keys() []string {
	users := idx.List()
	keys := make([]string, len(users))

	for i, u := range users {
		keys[i] = idx.key(u)
	}

	return keys
}

// String gibt eine menschenlesbare Listen-Darstellung aller Schlüssel im Index zurück.
func (idx *Index) String() string {
	return fmt.Sprintf("%v", idx.Keys())
}

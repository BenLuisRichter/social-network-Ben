package index

import (
	"fmt"
	"social-network/core/user"
	"social-network/view/element"
)

// Insert fügt einen neuen Benutzer in den Index ein.
func (idx *Index) Insert(e *user.User) {
	current := idx.root

	for !current.IsEmpty() {
		key := idx.key(e)
		currentKey := idx.key(current.User)

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
		currentKey := idx.key(current.User)

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

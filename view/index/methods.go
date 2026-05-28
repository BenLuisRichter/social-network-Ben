package index

import (
	"fmt"
	"social-network/core/user"
	"social-network/view/element"
)

// Insert fügt einen neuen Benutzer in den Index ein.
func (idx *Index) Insert(e *user.User) {
	// begin:solution
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
	// end:solution
}

// Find sucht einen Benutzer im Index anhand eines Schlüssels.
func (idx *Index) Find(key string) *element.Element {
	// begin:solution
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
	// end:solution

	return nil
}

// List liefert alle Benutzer im Index in sortierter Reihenfolge.
func (idx *Index) List() []*user.User {
	// begin:solution
	return idx.root.List()
	// end:solution
}

// Keys liefert eine Liste aller Schlüssel im Index in sortierter Reihenfolge.
func (idx *Index) Keys() []string {
	keys := []string{}

	// begin:solution
	for _, u := range idx.List() {
		keys = append(keys, idx.key(u))
	}
	// end:solution

	return keys
}

// String gibt eine menschenlesbare Listen-Darstellung aller Schlüssel im Index zurück.
func (idx *Index) String() string {
	// begin:solution
	return fmt.Sprintf("%v", idx.Keys())
	// end:solution
}

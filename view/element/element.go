package element

import "social-network/core/user"

// Element ist ein Element in einem binären Suchbaum, der Benutzer enthält.
type Element struct {
	user  *user.User
	left  *Element
	right *Element
}

// Empty erzeugt ein neues leeres Element.
func Empty() *Element {
	return &Element{
		user:  nil,
		left:  nil,
		right: nil,
	}
}

// IsEmpty prüft, ob dieses Element leer ist.
// Ein Element ist leer, wenn irgendeines der Felder nil ist.
func (e *Element) IsEmpty() bool {
	return e.user == nil || e.left == nil || e.right == nil
}

// SetUser setzt den Benutzer in diesem Element.
func (e *Element) SetUser(u *user.User) {
	if u == nil {
		return
	}

	e.user = u

	if e.IsEmpty() {
		e.left = Empty()
		e.right = Empty()
	}
}

// LeftChild gibt das linke Kind dieses Elements zurück.
func (e *Element) LeftChild() *Element {
	return e.left
}

// RightChild gibt das rechte Kind dieses Elements zurück.
func (e *Element) RightChild() *Element {
	return e.right
}

// User gibt den Benutzer zurück, der in diesem Element gespeichert ist.
func (e *Element) User() *user.User {
	return e.user
}

// List gibt eine Liste aller Benutzer in diesem Element und seinen Kindern
// in sortierter Reihenfolge zurück.
func (e *Element) List() []*user.User {
	if e.IsEmpty() {
		return []*user.User{}
	}

	users := []*user.User{}
	users = append(users, e.LeftChild().List()...)
	users = append(users, e.User())
	users = append(users, e.RightChild().List()...)
	return users
}

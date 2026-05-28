package element

import "social-network/core/user"

// Element ist ein Element in einem binären Suchbaum, der Benutzer enthält.
type Element struct {
	User  *user.User
	Left  *Element
	Right *Element
}

// Empty erzeugt ein neues leeres Element.
func Empty() *Element {
	return &Element{
		User:  nil,
		Left:  nil,
		Right: nil,
	}
}

// IsEmpty prüft, ob dieses Element leer ist.
// Ein Element ist leer, wenn irgendeines der Felder nil ist.
func (e *Element) IsEmpty() bool {
	// begin:solution
	return e.User == nil || e.Left == nil || e.Right == nil
	// end:solution
}

// SetUser setzt den Benutzer in diesem Element.
func (e *Element) SetUser(u *user.User) {
	// begin:solution
	if u == nil {
		return
	}

	e.User = u

	if e.IsEmpty() {
		e.Left = Empty()
		e.Right = Empty()
	}
	// end:solution
}

// List gibt eine Liste aller Benutzer in diesem Element und seinen Kindern
// in sortierter Reihenfolge zurück.
func (e *Element) List() []*user.User {
	users := []*user.User{}

	// begin:solution
	if e.IsEmpty() {
		return users
	}

	users = append(users, e.Left.List()...)
	users = append(users, e.User)
	users = append(users, e.Right.List()...)
	// end:solution

	return users
}

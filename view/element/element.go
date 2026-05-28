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
	// Hinweis:
	// - Prüfen Sie, ob `e.User`, `e.Left` oder `e.Right` nil ist. Wenn eines davon nil ist, ist dieses Element leer.

	return e.User == nil || e.Left == nil || e.Right == nil
}

// SetUser setzt den Benutzer in diesem Element.
func (e *Element) SetUser(u *user.User) {
	// Hinweis:
	// - Wenn `u` nil ist, soll dieses Element leer bleiben, also tun Sie nichts.
	// - Andernfalls setzen Sie das `User`-Feld dieses Elements auf `u`.
	// - Wenn dieses Element vorher leer war, müssen Sie auch die Kinder dieses Elements auf neue leere Elemente setzen.

	if u == nil {
		return
	}

	e.User = u

	if e.IsEmpty() {
		e.Left = Empty()
		e.Right = Empty()
	}
}

// List gibt eine Liste aller Benutzer in diesem Element und seinen Kindern
// in sortierter Reihenfolge zurück.
func (e *Element) List() []*user.User {
	// Hinweis:
	// - Wenn dieses Element leer ist, geben Sie eine leere Liste zurück.
	// - Andernfalls rufen Sie `List` rekursiv auf dem linken Kind auf, fügen Sie den Benutzer dieses Elements
	//   zur Liste hinzu und rufen Sie `List` rekursiv auf dem rechten Kind auf.
	// - Kombinieren Sie die Ergebnisse zu einer einzigen Liste.

	users := []*user.User{}

	if e.IsEmpty() {
		return users
	}

	users = append(users, e.Left.List()...)
	users = append(users, e.User)
	users = append(users, e.Right.List()...)

	return users
}

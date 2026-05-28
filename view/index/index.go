package index

import (
	"social-network/core/user"
	"social-network/view/element"
)

// Index repräsentiert einen binären Suchbaum, der Benutzer enthält.
// Der Baum dient als Suchindex, um Benutzer schnell anhand eines Schlüssels zu finden.
type Index struct {
	root *element.Element
	key  func(e *user.User) string
}

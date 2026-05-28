package user

// User repräsentiert einen Benutzer in unserem sozialen Netzwerk.
type User struct {
	Id       int
	Nickname string
	Name     string
	Surname  string
	Contacts []int
}

// New erstellt einen neuen Benutzer mit dem angegebenen Nickname
// sowie Vor- und Nachname. Die Kontaktliste ist zunächst leer.
func New(id int, nickname string, name string, surname string) *User {
	return &User{
		Id:       id,
		Nickname: nickname,
		Name:     name,
		Surname:  surname,
		Contacts: []int{},
	}
}

// AddContact fügt einen Kontakt zur Kontaktliste des Benutzers hinzu.
func (u *User) AddContact(contactID int) {
	u.Contacts = append(u.Contacts, contactID)
}

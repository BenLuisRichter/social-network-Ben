package database

import (
	"fmt"
	"social-network/core/user"
)

func ExampleDatabase_UsersByNickname() {
	db := testdata()

	fmt.Println(db.UsersByNickname())

	// Output:
	// [michkenntjeder superstar123 zombie43]
}

func ExampleDatabase_UsersByName() {
	db := testdata()

	fmt.Println(db.UsersByName())

	// Output:
	// [Arno Eva Max]
}

func ExampleDatabase_UsersBySurname() {
	db := testdata()

	fmt.Println(db.UsersBySurname())

	// Output:
	// [Einbildung Mustermann Nym]
}

func testdata() *Database {
	db := New()

	db.AddUser(user.New("michkenntjeder", "Max", "Mustermann"))
	db.AddUser(user.New("zombie43", "Arno", "Nym"))
	db.AddUser(user.New("superstar123", "Eva", "Einbildung"))

	return db
}

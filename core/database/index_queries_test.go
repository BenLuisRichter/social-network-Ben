package database

import (
	"fmt"
	"social-network/core/user"
)

func ExampleDatabase_UsersById() {
	db := testdata()
	users := db.UsersById()

	fmt.Println(users)

	// Output:
	// [1 2 3]
}

func ExampleDatabase_UsersByNickname() {
	db := testdata()
	users := db.UsersByNickname()

	fmt.Println(users)

	// Output:
	// [michkenntjeder superstar123 zombie43]
}

func ExampleDatabase_UsersByName() {
	db := testdata()
	users := db.UsersByName()

	fmt.Println(users)

	// Output:
	// [Arno Eva Max]
}

func ExampleDatabase_UsersBySurname() {
	db := testdata()
	users := db.UsersBySurname()

	fmt.Println(users)

	// Output:
	// [Einbildung Mustermann Nym]
}

func testdata() *Database {
	db := New()

	db.AddUser(user.New(1, "michkenntjeder", "Max", "Mustermann"))
	db.AddUser(user.New(2, "zombie43", "Arno", "Nym"))
	db.AddUser(user.New(3, "superstar123", "Eva", "Einbildung"))

	return db
}

package main

import "fmt"

type User struct {
	ID                         int
	FirstName, LastName, Email string
}

// YOUR CODE HERE

var user = User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn@gmail.com"}

func updateEmail(u *User) {
	u.Email = ""
	fmt.Println(u.Email)
}

func main() {
	updateEmail(&user)
	fmt.Println(user)
	fmt.Println("Pointers!")
}

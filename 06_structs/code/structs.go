package main

import "fmt"

// User is a user type
type User struct {
	ID                         int
	FirstName, LastName, Email string
}

type Group struct {
	role           string
	users          []User
	newestUser     User
	spaceAvailable bool
}

func describeUser(u User) string {
	desc := fmt.Sprintf("%s %s (%d)", u.FirstName, u.LastName, u.ID)
	return desc
}

func describeGroup(g Group) string {
	if len(g.users) > 2 {
		g.spaceAvailable = false
	}
	desc := fmt.Sprintf("%s group has %d members and the newest users is %s", g.role, len(g.users), g.newestUser.FirstName)
	return desc
}

// // User is a user type
// type User struct {
// ID					       int
// FirstName, LastName, Email string
// }

func main() {
	u := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}
	u2 := User{ID: 2, FirstName: "Marilen", LastName: "Monroe", Email: "marilyn.monrae@gmail.com"}
	u3 := User{ID: 3, FirstName: "Marilen", LastName: "Monroe", Email: "marilyn.monrae@gmail.com"}
	g := Group{
		role: "admin",
		users: []User{
			u,
			u2,
			u3,
		},
		newestUser:     u2,
		spaceAvailable: true,
	}
	fmt.Println(describeUser(u))
	fmt.Println(describeGroup(g))
}

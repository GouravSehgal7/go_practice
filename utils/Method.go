package utils

import "fmt"

// functions in classes are called method
// similarly function associated with struct are method

type User struct {
	Email string
	Name  string
	Age   int
}

func (u User) Getuserdet() {
	fmt.Println("user details are ", u.Email, u.Age, u.Name)
}

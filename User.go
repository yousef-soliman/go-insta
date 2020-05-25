package main

import "fmt"

type User struct {
	id   int
	name string
}

func (user User) String() string {
	return fmt.Sprintf("id = %d, name=%s", user.id, user.name)
}

type users struct {
	users []User
}

type UserInterface interface {
	init()
	createNewUser(User)
	showAll()
}

func (usersList *users) init() {
	usersList.users = make([]User, 0)
}
func (usersList *users) createNewUser(user User) {
	usersList.users = append(usersList.users, user)
}

func (usersList users) showAll() {
	for _, user := range usersList.users {
		fmt.Println(user)

	}
}

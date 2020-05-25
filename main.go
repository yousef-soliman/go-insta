package main

func main() {
	var userslist users
	userslist.init()
	var u = User{id: 1, name: "Yousef"}
	var u2 = User{id: 2, name: "Yousef"}
	userslist.createNewUser(u)
	userslist.createNewUser(u2)
	userslist.createNewUser(u2)
	userslist.showAll()
}

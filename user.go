package main

type user struct {
	name string
	age  int16
}

func createUser(name string, age int16) user {

	newUser := user{
		name: name,
		age:  age,
	}

	return newUser

}

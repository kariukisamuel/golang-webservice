package models

import (
	"errors"
	"fmt"
)

type User struct {
	Name     string
	Id       int
	Location string
}

var (
	users  []*User
	nextId int = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	if u.Id == 0 {
		return User{}, errors.New("new user must not include id or it must be equal to zero")
	}
	u.Id = nextId
	nextId++
	users = append(users, &u)
	return u, nil
}

func GetUserById(id int) (User, error) {
	for _, u := range users {
		if u.Id == id {
			return *u, nil
		}
	}

	return User{}, fmt.Errorf("User With Id %v Not Found", id)

}

func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.Id == u.Id {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User With Id %v Not Found", u.Id)
}

func RemoverUserById(id int) error {
	for i, u := range users {
		if u.Id == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User With Id %v Not Found", id)
}

package models

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

func AddUsers(u User) (User, error) {
	u.Id = nextId
	nextId++
	users = append(users, &u)
	return u, nil
}

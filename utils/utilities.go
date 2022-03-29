package utils

type User struct {
	Name string
}

func CreateUser() User {
	user := User{"Kobi"}
	return user
}

func (user *User) ChangeUserName(name string) {
	user.Name = name
}

func UpdateName(name *string) {
	*name = "Sobi"
}

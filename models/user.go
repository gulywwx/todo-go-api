package models

import (
	"errors"

	"github.com/gulywwx/todo-go-api/utils"
)

type User struct {
	ID       string `json:"id,omitempty" `
	Name     string `json:"name,omitempty" `
	Password string `json:"password,omitempty" `
	Desc     string `json:"desc,omitempty" `
}

var (
	userlist = make(map[string]*User)
)

func init() {
	userlist["admin"] = &User{ID: "1", Name: "admin", Password: "admin"}
}

func FindUser(username string) (*User, error) {
	user, found := userlist[username]

	if !found {
		return nil, errors.New("user not found")
	}

	return user, nil
}

/**
创建用户账号
*/
func CreateUser(username string, password string) (*User, error) {
	user := User{
		ID:       utils.GenUlid(),
		Name:     username,
		Password: password,
	}
	userlist[username] = &user

	return &user, nil
}

func ResetPass(username string, oldpassword string, newpassword string) error {
	user, found := userlist[username]

	if !found {
		return errors.New("user not found")
	}
	if oldpassword != user.Password {
		return errors.New("old pass is incorrect")
	}

	user.Password = newpassword

	// log.Println(user)

	return nil
}

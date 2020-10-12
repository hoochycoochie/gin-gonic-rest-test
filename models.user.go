package main

import (
	"errors"
	"strings"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

var userList []user = []user{
	user{Username: "user1", Password: "pass1"},
	user{Username: "user2", Password: "pass2"},
	user{Username: "user3", Password: "pass3"},
}

func registerNewUser(username, password string) (*user, error) {

	if strings.TrimSpace(password) == "" {
		return nil, errors.New("password empty")
	} else if strings.TrimSpace(username) == "" {
		return nil, errors.New("username empty")
	} else if !isUsernameAvailable(username) {
		return nil, errors.New("username not available")
	}
	u := user{Username: username, Password: password}

	userList = append(userList, u)
	return &u, nil
}

func isUsernameAvailable(username string) bool {

	for _, v := range userList {
		if v.Username == username {
			return false
		}
	}
	return true
}

func isValidUser(username, password string) bool {

	for _, v := range userList {
		if v.Username == username && v.Password == password {
			return true
		}
	}
	return false
}

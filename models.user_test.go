package main

import "testing"

func TestUsernameAvailability(t *testing.T) {
	saveLists()
	if isUsernameAvailable("user1") {
		t.Fail()
	}

	if !isUsernameAvailable("toto") {
		t.Fail()
	}

	registerNewUser("newuser", "newpass")

	if isUsernameAvailable("newuser") {
		t.Fail()
	}

	restoreLists()
}

func TestValidUserRegistration(t *testing.T) {
	saveLists()

	u, err := registerNewUser("toto1", "toto")

	if err != nil || u.Username == "" {
		t.Fail()
	}

	restoreLists()
}

func TestInvalidUserRegistration(t *testing.T) {
	saveLists()

	u, err := registerNewUser("user1", "pass1")

	if err == nil || u != nil {
		t.Fail()
	}

	u, err = registerNewUser("newuser", "")

	if err == nil || u != nil {
		t.Fail()
	}

	restoreLists()
}

func TestUserValidity(t *testing.T) {

	if !isValidUser("user1", "pass1") {
		t.Fail()
	}

	if isValidUser("&&&&&", "&&&") {
		t.Fail()
	}

	if isValidUser("", "pass1") {
		t.Fail()
	}
}

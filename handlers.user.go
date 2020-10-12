package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func generateSessionToken() string {
	return strconv.FormatInt(rand.Int63(), 16)
}
func showRegistrationPage(c *gin.Context) {
	render(c, gin.H{"title": "Register"}, "register.html")
}
func showLoginPage(c *gin.Context) {

}

func performLogin(c *gin.Context) {

}

func logout(c *gin.Context) {

}

func register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if _, err := registerNewUser(username, password); err == nil {
		render(c, gin.H{
			"title": "Successful registration & Login",
		}, "login-successful.html")
	} else {

		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"ErrorTitle":   "Registration Failed",
			"ErrorMessage": err.Error()})
	}

}

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
	render(c, gin.H{"title": "Login"}, "login.html")
}

func performLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if isValidUser(username, password) {

		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)

		c.HTML(http.StatusOK, "login-successful.html", gin.H{
			"title": "Successful login",
		})
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"ErrorTitle":   "Login failed",
			"ErrorMessage": "Invalid credentials provided",
		})
	}

}

func logout(c *gin.Context) {

	c.SetCookie("token", "", -1, "", "", false, true)
	c.Redirect(http.StatusTemporaryRedirect, "/")

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

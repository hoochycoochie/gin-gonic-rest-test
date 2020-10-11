package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	router = gin.Default()
	router.LoadHTMLGlob("templates/*")

	initializeRoutes()

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}

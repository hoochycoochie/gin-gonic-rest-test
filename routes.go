package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := GetAllArticles()
	//	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Home Page", "payload": articles})

	render(c, gin.H{"title": "Home Page", "payload": articles}, "index.html")
}

func getArticle(c *gin.Context) {
	// Check if the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Check if the article exists
		if article, err := getArticleByID(articleID); err == nil {
			// Call the HTML method of the Context to render a template
			/* c.HTML(
				// Set the HTTP status to 200 (OK)
				http.StatusOK,
				// Use the index.html template
				"article.html",
				// Pass the data that the page uses
				gin.H{
					"title":   article.Title,
					"payload": article,
				},
			) */
			render(c, gin.H{
				"title":   article.Title,
				"payload": article,
			}, "article.html")

		} else {
			// If the article is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}

}
func initializeRoutes() {
	router.GET("/", showIndexPage)

	userRoutes := router.Group("/u")
	{
		userRoutes.GET("/register", showRegistrationPage)
		userRoutes.POST("/register", register)
	}
	router.GET("/article/view/:article_id", getArticle)
}

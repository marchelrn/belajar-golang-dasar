package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func multipartFormRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		marchel := c.DefaultPostForm("Marchel", "Anonymous")

		c.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": message,
			"orn":     marchel,
		})
	})
	return router
}

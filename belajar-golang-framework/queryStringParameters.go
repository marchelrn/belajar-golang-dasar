package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupQueryStringParameters() *gin.Engine {
	router := gin.Default()

	router.GET("/welcome", func(c *gin.Context) {
		firstName := c.DefaultQuery("firstName", "Guest")
		lastName := c.Query("lastName")

		c.String(http.StatusOK, "Hello %s %s \n", firstName, lastName)
	})

	return router
}

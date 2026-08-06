package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HTTP Method

func getting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Method": "GET"})
}

func posting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Method": "POST"})
}

func putting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Method": "PUT"})
}

func deleting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Method": "DELETE"})
}

func patching(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Method": "PATCH"})
}

func head(c *gin.Context) {
	c.Status(http.StatusOK)
}

func options(c *gin.Context) {
	c.Status(http.StatusOK)
}

func setupHttpMethodRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	httpMethod := router.Group("/httpMethod")

	httpMethod.GET("/someGet", getting)
	httpMethod.POST("/somePost", posting)
	httpMethod.PUT("/somePut", putting)
	httpMethod.PATCH("/somePatch", patching)
	httpMethod.DELETE("/someDelete", deleting)
	httpMethod.HEAD("/someHead", head)
	httpMethod.OPTIONS("/someOptions", options)

	return router
}

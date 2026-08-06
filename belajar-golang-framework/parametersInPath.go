package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupParametersInPathRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/pong", func(c *gin.Context) {
		c.String(http.StatusOK, "ping \n")
	})

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s \n", name)
	})

	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action + "\n"

		c.String(http.StatusOK, message)
	})

	return router
}

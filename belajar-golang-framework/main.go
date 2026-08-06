package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// servers
	gin.SetMode(gin.DebugMode)

	httpMethod := setupHttpMethodRoutes()
	parametersInPath := setupParametersInPathRoutes()
	queryStringParam := setupQueryStringParametersRoutes()
	routing := setupRouting()
	multiPart := multipartFormRoutes()

	m := map[int]*gin.Engine{
		1: httpMethod,
		2: parametersInPath,
		3: queryStringParam,
		4: routing,
		5: multiPart,
	}
	srv := &http.Server{
		Addr:    ":8000",
		Handler: m[5],
	}

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Failed to start server: %v", err)
	}

}

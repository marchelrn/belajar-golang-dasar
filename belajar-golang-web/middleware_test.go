package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before execute Handler")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After execute Handler")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic:", err)
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error: %v", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprint(writer, "Hello Middleware")
	})

	mux.HandleFunc("/foo", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("foo Executed")
		fmt.Fprint(writer, "Hello Foo")
	})

	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Panic Executed")
		panic("ups")
	})

	logMiddleware := LogMiddleware{
		Handler: mux,
	}

	errorHandler := ErrorHandler{
		Handler: &logMiddleware,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: &errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

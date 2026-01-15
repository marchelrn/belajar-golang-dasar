package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {

	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		// logic web
		_, err := fmt.Fprint(writer, "Hello World")
		if err != nil {
			panic(err)
		}
	}

	server := http.Server{
		Addr:    "localhost:9090",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		t.Fatal(err)
	}

}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Hello World")
		if err != nil {
			panic(err)
		}
	})

	mux.HandleFunc("/about", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "About Page")
		if err != nil {
			panic(err)
		}
	})
	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Images")
		if err != nil {
			panic(err)
		}
	})
	mux.HandleFunc("/images/thumbnails/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Thumbnails Images")
		if err != nil {
			panic(err)
		}
	})
	server := http.Server{
		Addr:    "localhost:9090",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		t.Fatal(err)
	}
}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprintln(writer, request.Method)
		if err != nil {
			t.Fatal(err)
		}
		_, err = fmt.Fprintln(writer, request.RequestURI)
		if err != nil {
			t.Fatal(err)
		}
	}

	server := http.Server{
		Addr:    "localhost:9090",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		t.Fatal(err)
	}
}

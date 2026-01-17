package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	ContentType := request.Header.Get("Content-Type")
	fmt.Fprint(writer, "Content-Type: ", ContentType, "\r\n")
}

func TestRequestHeader(t *testing.T) {
	Request := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
	Request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	RequestHeader(recorder, Request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func ResponseHeader(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")
	writer.Header().Add("X-Powered-By", "Golang")
	_, err := fmt.Fprint(writer, "OK")
	if err != nil {
		panic(err)
	}
}

func TestResponseHeader(t *testing.T) {
	Request := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
	Request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	ResponseHeader(recorder, Request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

	fmt.Println("Content-Type:", response.Header.Get("Content-Type"))
	fmt.Println("X-Powered-By:", response.Header.Get("X-Powered-By"))
}

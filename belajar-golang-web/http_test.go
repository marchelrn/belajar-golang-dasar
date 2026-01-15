package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HalloHandler(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintln(writer, "Hello World")
	if err != nil {
		panic(err)
	}
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9090/hello", nil)
	recorder := httptest.NewRecorder()

	HalloHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}

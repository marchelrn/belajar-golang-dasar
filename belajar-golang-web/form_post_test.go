package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}
	firstname := request.PostForm.Get("firstname")
	lastname := request.PostForm.Get("lastname")
	_, err = fmt.Fprintf(writer, "Hello %s %s", firstname, lastname)
	if err != nil {
		panic(err)
	}
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("firstname=Marchel&lastname=Manullang")
	request := httptest.NewRequest(http.MethodPost, "/", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

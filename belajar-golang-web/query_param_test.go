package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		_, err := fmt.Fprint(writer, "Hello")
		if err != nil {
			panic(err)
		}
	} else {
		_, err := fmt.Fprintf(writer, "Hello %s", name)
		if err != nil {
			panic(err)
		}
	}
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9090/hello?name=", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}

func MultipleQueryParameter(writer http.ResponseWriter, request *http.Request) {
	firstName := request.URL.Query().Get("first_name")
	lastName := request.URL.Query().Get("last_name")

	_, err := fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
	if err != nil {
		panic(err)
	}
}

func TestMultipleQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9090/hello?first_name=Marchel&last_name=Manullang", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParameter(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}

func MultipleParameterValue(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	names := query["name"]

	_, err := fmt.Fprint(writer, "Hello ", strings.Join(names, " "))
	if err != nil {
		panic(err)
	}
}

func TestMultipleParameterValue(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9090/hello?name=Marchel&name=Manullang", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValue(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}

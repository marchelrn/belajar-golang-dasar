package belajar_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateDataMap(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("templates/name.gohtml"))
	err := t.ExecuteTemplate(writer, "name.gohtml", map[string]interface{}{
		"title": "Marchel's Page",
		"name":  "Marchel",
	})

	if err != nil {
		panic(err)
	}
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	bodyBytes, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyBytes))
}

// Struct

type Page struct {
	title string
	name  string
}

func TemplateDataStruct(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("templates/name.gohtml"))
	err := t.ExecuteTemplate(writer, "name.gohtml", Page{
		title: "Marchel's Page",
		name:  "Marchel",
	})
	if err != nil {
		panic(err)
	}
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	bodyBytes, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyBytes))
}

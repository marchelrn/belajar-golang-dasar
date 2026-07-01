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
		"Title": "Marchel's Page",
		"Name":  "Marchel",
		"Address": map[string]interface{}{
			"Street": "Jalan belum ada",
		},
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
	Title   string
	Name    string
	Address Address
}

type Address struct {
	Street string
}

func TemplateDataStruct(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("templates/name.gohtml"))
	err := t.ExecuteTemplate(writer, "name.gohtml", Page{
		Title: "Marchel's Page",
		Name:  "Marchel",
		Address: Address{
			Street: "Malalayang",
		},
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

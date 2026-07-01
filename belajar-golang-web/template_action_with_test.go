package belajar_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Data struct {
	Title   string
	Name    string
	Address any
}

type Addresses struct {
	Title, Street, City, Name string
}

func TemplateActionWith(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("templates/with.gohtml"))
	err := t.ExecuteTemplate(writer, "with.gohtml", Data{
		Title: "Go With",
		Name:  "Marchel",
		//Address: Addresses{
		//	Title:  "Go With",
		//	Name:   "Marchel",
		//	Street: "Jln 5 september",
		//	City:   "Manado",
		//},
	})

	if err != nil {
		panic(err)
	}
}

func TestTemplateActionWith(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)

	TemplateActionWith(recorder, request)

	bodyBytes, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyBytes))
}

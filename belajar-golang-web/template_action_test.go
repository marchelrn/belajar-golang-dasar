package belajar_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateIf(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("templates/condition.gohtml"))
	err := t.ExecuteTemplate(writer, "condition.gohtml", map[string]interface{}{
		"Title":    "Condition",
		"Name":     "Marchel",
		"LastName": "Manullang",
	})

	if err != nil {
		panic(err)
	}
}

func TestTemplateIf(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)

	TemplateIf(recorder, request)

	bodyBytes, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyBytes))
}

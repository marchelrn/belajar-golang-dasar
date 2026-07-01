package belajar_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateActionRange(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("templates/range.gohtml"))
	err := t.ExecuteTemplate(writer, "range.gohtml", map[string]interface{}{
		"Title": "Range",
		"Hobbies": []string{
			"Game", "Coding", "Music",
		},
	})

	if err != nil {
		panic(err)
	}
}

func TestTemplateActionRange(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)

	TemplateActionRange(recorder, request)

	bodyBytes, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyBytes))
}

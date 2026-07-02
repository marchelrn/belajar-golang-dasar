package belajar_golang_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed templates/*.gohtml
var templates embed.FS

var myTemplates = template.Must(template.ParseFS(templates, "templates/*.gohtml"))

func TemplateCaching(writer http.ResponseWriter, request *http.Request) {
	err := myTemplates.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template")
	if err != nil {
		panic(err)
	}
}

func TemplateCachingName(writer http.ResponseWriter, request *http.Request) {
	err := myTemplates.ExecuteTemplate(writer, "name.gohtml", map[string]interface{}{
		"Title": "testing from caching",
		"Name":  "Marchel",
		"Address": map[string]interface{}{
			"Street": "123 Street",
		},
	})
	if err != nil {
		panic(err)
	}
}

func TestTemplateCaching(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)

	TemplateCachingName(recorder, request)

	bodyBytes, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyBytes))
}

package belajar_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateComparator(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("templates/comparator.gohtml"))
	err := t.ExecuteTemplate(writer, "comparator.gohtml", map[string]interface{}{
		"Title":      "Comparator",
		"FinalValue": 59,
	})
	if err != nil {
		panic(err)
	}
}

func TestTemplateComparator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateComparator(recorder, request)

	bodyBytes, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyBytes))
}

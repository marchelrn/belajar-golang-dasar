package belajar_golang_web

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func UploadForm(writer http.ResponseWriter, request *http.Request) {
	err := myTemplates.ExecuteTemplate(writer, "upload_form.gohtml", nil)
	if err != nil {
		panic(err)
	}
}

func Upload(writer http.ResponseWriter, request *http.Request) {
	file, fileHeader, err := request.FormFile("file")
	if err != nil {
		panic(err)
	}
	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}
	name := request.PostFormValue("name")
	err = myTemplates.ExecuteTemplate(writer, "upload_success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
	if err != nil {
		panic(err)
	}
}

func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {

	}
}

//go:embed "resources/letter_m.png"
var UploadFileTest []byte

func TestUploadFile(t *testing.T) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	err := writer.WriteField("name", "Marchel Manullang")
	if err != nil {
		return
	}
	file, _ := writer.CreateFormFile("file", "CONTOHUPLOAD.png")
	_, err = file.Write(UploadFileTest)
	if err != nil {
		panic(err)
	}
	err = writer.Close()
	if err != nil {
		return
	}
	request := httptest.NewRequest(http.MethodPost, "localhost:8080/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	bodyBytes, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyBytes))
}

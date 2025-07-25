package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SimpleHtmlTemplate(writer http.ResponseWriter, request *http.Request) {
	templateText := `<html><body>{{.}}<body></html>`

	// Versi 1 dengan menampilkan baris kode err
	// tmpl, err := template.New("SIMPLE").Parse(templateText)
	// if err != nil {
	// 	panic(err)
	// }

	// versi 2 dengan Must(): tidak menampilkan baris kode err
	tmpl := template.Must(template.New("SIMPLE").Parse(templateText))
	tmpl.ExecuteTemplate(writer, "SIMPLE", "Hello World Template")
}

func TestSimpleHtmlTemplate(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHtmlTemplate(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}


func SimpleHtmlTemplateFile(writer http.ResponseWriter, request *http.Request){
	tmpl := template.Must(template.ParseFiles("./templates/simple.gohtml"))
	tmpl.ExecuteTemplate(writer, "simple.gohtml", "Hello World Template")
}

func TestSimpleHtmlTemplateFile(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHtmlTemplateFile(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}


func SimpleHtmlTemplateDirectory(writer http.ResponseWriter, request *http.Request){
	tmpl := template.Must(template.ParseGlob("./templates/*.gohtml"))
	tmpl.ExecuteTemplate(writer, "simple.gohtml", "Hello World Template")
}

func TestSimpleHtmlTemplateDirectory(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHtmlTemplateDirectory(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// kode dipindahkan ke template_caching_test.go

func SimpleHtmlTemplateEmbed(writer http.ResponseWriter, request *http.Request){
	tmpl := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	tmpl.ExecuteTemplate(writer, "simple.gohtml", "Hello World Template")
} 

func TestSimpleHtmlTemplateEmbed(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHtmlTemplateEmbed(recorder, request)	

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateDataMap(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/name.gohtml"))
	tmpl.ExecuteTemplate(writer, "name.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Name":  "Lev",
	})
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
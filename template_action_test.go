package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateActionIf(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/if.gohtml"))
	tmpl.ExecuteTemplate(writer, "if.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Admin": true, 
		// "Member": true, 
		"Name":  "Lev",
	})
}

func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}


// ? OPERATOR PERBANDINGAN GO-LANG
// eq / equal						==>		== 	(sama dengan)
// ne / not equal				==>		!= 	(tidak sama dengan)
// lt / less than				==>	 	< 	(kurang dari)
// le / less equal			==> 	<= 	(lebih besar dari)
// gt / greater than		==> 	< 	(kurang dari sama dengan)
// ge / greater equal		==> 	>= 	(lebih besar dari sama dengan)

func TemplateActionOperator(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/comparator.gohtml"))
	tmpl.ExecuteTemplate(writer, "comparator.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"FinalValue":  85,
	})
}

func TestTemplateActionOperator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateActionOperator(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}


func TemplateActionRange(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/range.gohtml"))
	tmpl.ExecuteTemplate(writer, "range.gohtml", map[string]interface{}{
		"Title": "Template Action Range",
		"Hobbies":  []string{
			// "Coding", "Gaming", "Sleeping",
		},
	})
}

func TestTemplateActionRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
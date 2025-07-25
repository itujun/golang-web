package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type myInfo struct {
	Name string
}

func (info myInfo) SayHello(targetName string) string {
	return "Hello " + targetName + ", my name is " + info.Name
}

func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.New("FUNCTION").Parse(`{{ .SayHello "Lev" }}`))

	tmpl.ExecuteTemplate(writer, "FUNCTION", myInfo{Name: "Juna"})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateFunctionGlobal(writer http.ResponseWriter, request *http.Request) {
	// ? Link Global Function: https://github.com/golang/go/blob/master/src/text/template/funcs.go
	tmpl := template.Must(template.New("FUNCTION").Parse(`{{ len .Name}}`))

	tmpl.ExecuteTemplate(writer, "FUNCTION", myInfo{Name: "Learn Go-Lang"})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}


func TemplateFunctionGlobalCreate(writer http.ResponseWriter, request *http.Request) {
	// membuat function baru
	tmpl := template.New("FUNCTION")

	// daftarkan global function yang sudah dibuat
	tmpl = tmpl.Funcs(map[string]interface{}{
		"upper": func(val string) string {
			return strings.ToUpper(val)
		},
	})

	// parse template
	tmpl = template.Must(tmpl.Parse(`{{ upper .Name }}`))

	tmpl.ExecuteTemplate(writer, "FUNCTION", myInfo{Name: "Lev Tempest"})
}

func TestTemplateFunctionGlobalCreate(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobalCreate(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}


func TemplateFunctionPipeline(writer http.ResponseWriter, request *http.Request) {
	// membuat function baru
	tmpl := template.New("FUNCTION")

	// daftarkan global function yang sudah dibuat
	tmpl = tmpl.Funcs(map[string]interface{}{
		"sayHello": func(name string) string {
			return "Hello " + name
		},
		"upper": func(val string) string {
			return strings.ToUpper(val)
		},
	})

	// parse template
	tmpl = template.Must(tmpl.Parse(`{{ sayHello .Name | upper }}`))

	tmpl.ExecuteTemplate(writer, "FUNCTION", myInfo{Name: "Lev Tempest"})
}

func TestTemplateFunctionPipeline(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionPipeline(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
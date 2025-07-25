package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
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
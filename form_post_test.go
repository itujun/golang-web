package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {
	// PARSING FORM POST MANUAL
	// err := request.ParseForm()
	// if err != nil {
	// 	panic(err)
	// }

	// firstName := request.PostForm.Get("first_name")
	// lastName := request.PostForm.Get("last_name")

	// fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
	// END PARSING FORM POST MANUAL

	// PARSING OTOMATIS FORM POST GOLANG
	firstName := request.PostFormValue("first_name")
	lastName := request.PostFormValue("last_name")
	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
	// END PARSING OTOMATIS FORM POST GOLANG
}

func TestFormPost(t *testing.T){
	requestBody := strings.NewReader("first_name=Lev&last_name=Tempest")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded") // wajib untuk form post
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
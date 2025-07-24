package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
	fmt.Fprintln(w, r.Method)
	fmt.Fprintln(w, r.URL)
}

func TestHelloHandler(t *testing.T) {
	// request := httptest.NewRequest("GET", "http://localhost:8080/", nil) // ini juga bisa
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	// CEK HASIL TEST
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}
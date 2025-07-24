package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(w http.ResponseWriter, r *http.Request){
	name := r.URL.Query().Get("name") // get parameter name pada url
	if name == "" {
		fmt.Fprint(w, "Hello")
	}else{
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/?name=Lev", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	
	fmt.Println(string(body))
}


func SayHelloMultiple(w http.ResponseWriter, r *http.Request){
	firstname := r.URL.Query().Get("first_name") // get parameter name pada url
	lastname := r.URL.Query().Get("last_name") // get parameter name pada url
	
	fmt.Fprintf(w, "Hello %s %s", firstname, lastname)
}

func TestMultipleQueryParameter(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/?first_name=Lev&last_name=Tempest", nil)
	recorder := httptest.NewRecorder()

	SayHelloMultiple(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	
	fmt.Println(string(body))
}


func SayHelloMultipleValue(w http.ResponseWriter, r *http.Request){
	query := r.URL.Query()
	names := query["name"]
	
	fmt.Fprintf(w, "Hello %s", strings.Join(names, " "))
}
func TestMultipleValueQueryParameter(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/?name=Lev&name=Tempest&name=san", nil)
	recorder := httptest.NewRecorder()

	SayHelloMultipleValue(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	
	fmt.Println(string(body))
}
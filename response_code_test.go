package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name") // get parameter name pada url
	if name == "" {
		writer.WriteHeader(http.StatusBadRequest)
		// writer.WriteHeader(400) // sama dengan writer.WriteHeader(http.StatusBadRequest) 
		fmt.Fprint(writer, "name is empty")
	} else {
		// writer.WriteHeader(http.StatusOK) // tidak perlu ditambahkan tidak apa-apa karena defaultnya adalah http.StatusOK / 200
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestResponseCode(t *testing.T){
	// request := httptest.NewRequest("GET", "http://localhost:8080/", nil) // bad request 400
	request := httptest.NewRequest("GET", "http://localhost:8080/?name=Lev", nil) // ok 200
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}
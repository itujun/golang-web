package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MEMBUAT COOKIE
func SetCookie(writer http.ResponseWriter, request *http.Request) {
	// cookie := http.Cookie{Name: "name", Value: "lev", Path: "/", HttpOnly: true} // ini juga bisa
	cookie := new(http.Cookie)
	cookie.Name = "Theme"
	cookie.Value = request.URL.Query().Get("theme")
	cookie.Path = "/"

	http.SetCookie(writer, cookie)
	fmt.Fprint(writer, "Success Set Cookie.")
}

// MENGAMBIL COOKIE
func GetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("Theme")
	if err != nil {
		fmt.Fprint(writer, "No Cookie")
	} else {
		fmt.Fprintf(writer, "Theme %s", cookie.Value)
	}
}

// TEST SET COOKIE
func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?theme=dark", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	
	cookies := recorder.Result().Cookies()
	for _, cookie := range cookies{
		fmt.Printf("Cookie %s:%s \n", cookie.Name, cookie.Value)
	}
}

// TEST GET COOKIE
func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	cookie := new(http.Cookie)
	cookie.Name = "Theme"
	cookie.Value = "dark"
	request.AddCookie(cookie)

	GetCookie(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
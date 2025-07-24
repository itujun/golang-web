package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		// ? w/writer di parameter ialah response yg akan ditampilkan pada client
		// ? r/request di parameter ialah request yg dikirimkan oleh client

		// logic web
		fmt.Fprint(w, "Hello World")
	}

	server := http.Server{Addr: "localhost:8080", Handler: handler}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
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

func TestMux(t *testing.T){
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World")
	})
	mux.HandleFunc("/product", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Daftar Produk")
	})
	mux.HandleFunc("/product/laptop/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Info Produk Laptop")
	})
	
	server := http.Server{Addr: "localhost:8080", Handler: mux}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
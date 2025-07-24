package golangweb

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	// tentukan direktori "yg akan digunakan sebagai handler
	directory := http.Dir("./resources")
	// buat file server dari direktori 
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer)) // http.StripPrefix() untuk menghilangkan prefix /static/ dari url

	server := http.Server{Addr: "localhost:8080", Handler: mux}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}


//go:embed resources
var resources embed.FS

func TestFileServerEmbed(t *testing.T) {
	directory, _ := fs.Sub(resources, "resources") // fs.Sub() untuk mengambil sub direktori
	// buat file server dari resources
	fileServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer)) // http.StripPrefix() untuk menghilangkan prefix /static/ dari url

	server := http.Server{Addr: "localhost:8080", Handler: mux}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
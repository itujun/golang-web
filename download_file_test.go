package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func DonwloadFile(writer http.ResponseWriter, request *http.Request) {
	file := request.URL.Query().Get("file")

	if file == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "Bad Request")
		return
	}

	// contoh tanpa content-disposition (gambar akan otomatis dirender oleh browser)
	// http.ServeFile(writer, request, "./resources/" + file)

	// meggunakan content-disposition (gambar tidak akan dirender, namun akan langsung di-download oleh user)
	writer.Header().Add("Content-Disposition", "attachment; filename=\"" + file + "\"")
	http.ServeFile(writer, request, "./resources/" + file)
}

func TestDownloadFile(t *testing.T){
	server := http.Server{Addr: "localhost:8080", Handler: http.HandlerFunc(DonwloadFile)}

	// ? jalankan server, lalu ketikkan pada url browser: // http://localhost:8080/?file=favicon.ico

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
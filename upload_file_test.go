package golangweb

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func UploadForm(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "upload.form.gohtml", nil)
}

// Upload handler
func Upload(writer http.ResponseWriter, request *http.Request) {
	// ambil file dari form
	file, fileHeader, err := request.FormFile("file") // FormFile() sudah otomatis (bawaan) melakukan parsing multipart 

	if(err != nil) {
		panic(err)
	}
	defer file.Close()

	// simpan file ke disk
	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if(err != nil) {
		panic(err)
	}
	defer fileDestination.Close()

	// copy file ke disk
	_, err = io.Copy(fileDestination, file)
	if(err != nil) {
		// misalnya error jika tidak memiliki permission untuk menulis ke disk
		panic(err)
	}

	// ambil yang bukan file pada form
	name := request.PostFormValue("name")
	myTemplates.ExecuteTemplate(writer, "upload.success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

func TestUploadFormServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/form", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./resources")))) // http.StripPrefix() untuk menghilangkan prefix /static/ dari url", Upload)

	server := http.Server{Addr: "localhost:8080", Handler: mux}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/favicon.ico
var uploadFileTest []byte

// Unit Test Upload File
func TestUpload(t *testing.T) {
	// buat body untuk menyimpan
	body := new(bytes.Buffer)

	// buat multipart writer untuk menulis body
	writer := multipart.NewWriter(body)

	writer.WriteField("name", "Lev Tempest") // field name
	file, _ :=writer.CreateFormFile("file", "gambar-lev-tempest.jpg") // field file

	// buat file
	file.Write(uploadFileTest)

	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType()) // jangan lupa set content type, jika tidak akan error
	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	bodyResponse, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyResponse))
}
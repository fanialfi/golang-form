package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func HandleProcess(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "", http.StatusBadRequest)
		return
	}

	// ParseMultipartForm digunakan untuk memparsing form data yang ada file nya
	// argument 1024 pada function tersebut adalah maxMemori
	// membuat file yang terupload akan disimpan sementara pada memori dengan alokasi adalah sebesar maxMemori
	// jika kapasitas file lebih besar, maka file akan disimpan didalam temporary file
	if err := req.ParseMultipartForm(1024); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// FormValue digunakan untuk pengambilan data dari view,
	// dimana key nya adalah name dari input form dari view
	alias := req.FormValue("alias")

	// FormFile digunakan untuk mengambil file yang diupload dan mengembalikan 3 object
	uploadFile, header, err := req.FormFile("file")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	defer uploadFile.Close()

	// os.Getwd mengembalikan path dari root ke current directory
	dir, err := os.Getwd()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	fileName := header.Filename

	if alias != "" {
		// filepath.Ext digunakan untuk mengambil ekstensi dari sebuah file (parameter) (example : .jpg, .png)
		fileName = fmt.Sprintf("%s%s", alias, filepath.Ext(header.Filename))
	}

	fileLocation := filepath.Join(dir, "files", fileName)

	// os.O_WRONLY, file yang dibukan hanya bisa digunakan untuk menulis saja
	// os.O_CREATE, jika file belum ada maka akan dibuat
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	defer targetFile.Close()

	// io.Copy mengisi konten dari parameter kedua ke parameter pertama
	if _, err := io.Copy(targetFile, uploadFile); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Write([]byte("done"))
}

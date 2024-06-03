package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// html에서 요청하는 css와 js파일 serving을 위한 핸들러
func serve_files(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/files/")

	//xxx.xxx.xxx/files/ 로 들어온 파일요청을 본 프로그램의 /webhandler 하위로 처리한다.
	http.ServeFile(w, r, "./websource/"+path)
}

func uploadForm(w http.ResponseWriter, r *http.Request) {
	tmpl := "./websource/index.html"
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fileName := r.FormValue("fileName")
	offset, err := strconv.ParseInt(r.FormValue("offset"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid offset", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("chunk")
	if err != nil {
		http.Error(w, "Error retrieving chunk", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// 해당 위치 디렉토리 없을 시 생성
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		os.Mkdir(uploadPath, os.ModePerm)
	}

	dst, err := os.OpenFile(uploadPath+fileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// 올바른 오프셋으로 위치
	_, err = dst.Seek(offset, io.SeekStart)
	if err != nil {
		http.Error(w, "Error seeking file", http.StatusInternalServerError)
		return
	}

	// 파일에 청크를 쓰기
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Error writing chunk", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Chunk at offset %d uploaded successfully", offset)
}

package main

import (
	"flag"
	"fmt"
	"net/http"
)

const uploadPath = "./uploads/"

func main() {
	portNumber := flag.String("port", "8080", "--port=<port number>")
	flag.Parse()

	http.HandleFunc("/", uploadForm)
	http.HandleFunc("/upload", uploadFile)
	http.HandleFunc("/files/", serve_files)
	fmt.Println("Starting server at :" + *portNumber)
	http.ListenAndServe("0.0.0.0:"+*portNumber, nil)
}

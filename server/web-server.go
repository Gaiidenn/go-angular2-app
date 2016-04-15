package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func loadIndex(w *http.ResponseWriter) {
	fmt.Println("loading index page")
	pwd, _ := os.Getwd()
	filepath := pwd + "/../client/index.html"
	index, _ := ioutil.ReadFile(filepath)

	io.WriteString(*w, string(index))
}

func serveFile(w *http.ResponseWriter, filePath string) {
	fmt.Println("page requested by the client : " + filePath)
	pwd, _ := os.Getwd()
	filePath = pwd + "/../client" + filePath
	fmt.Println(filePath)
	file, _ := ioutil.ReadFile(filePath)

	io.WriteString(*w, string(file))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New client request : ")
	if r.URL.Path == "/" {
		loadIndex(&w)
	} else {
		serveFile(&w, r.URL.Path)
	}
}

func main() {
	fmt.Println("Initializing handler...")
	http.HandleFunc("/", handler)
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

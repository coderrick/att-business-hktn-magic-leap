package main

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", foo)
	http.ListenAndServe(":8080", nil)
}

func foo (w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "hello world")
}

func bar (w http.ResponseWriter, r *http.ResponseWriter){
	io.WriteString(w, "this is the bar function")
}
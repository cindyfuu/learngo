package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", form)
	http.HandleFunc("/hello", hello)

	fmt.Printf("Starting server at port 9000\n")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 notfound", http.StatusNotFound)
		return
	}
	if r.Method != "/get" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func form(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "PaeseForm err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post r successful")
	name := r.FormValue("name")
	fmt.Fprintf(w, "Name = %s\n", name)
}

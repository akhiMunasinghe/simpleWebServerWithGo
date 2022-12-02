package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(responseWriter, "404 not found", http.StatusNotFound)
		return
	}
	if request.Method != "GET" {
		http.Error(responseWriter, "Method is not supported", http.StatusAccepted)
		return
	}
	fmt.Fprintf(responseWriter, "Hello!")
}

func formHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(responseWriter, "Parseform() err: %v", err)
		return
	}
	fmt.Fprintf(responseWriter, "POST request successful\n")
	name := request.FormValue("name")
	address := request.FormValue("address")
	fmt.Fprintf(responseWriter, "Name is %s\n", name)
	fmt.Fprintf(responseWriter, "Address is %s\n", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

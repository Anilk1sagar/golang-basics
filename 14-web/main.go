package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<h1>Hello Api</h1>")
}

func main() {

	fmt.Println("hello World!")

	http.HandleFunc("/", index)
	fmt.Println("Server Started on port: 9110")
	http.ListenAndServe(":9110", nil)
}

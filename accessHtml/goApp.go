package main

import (
	"fmt"
	"net/http"
)

func handler_func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go is cool")
}

func handler_func_about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go about is cool")
}

func main() {
	http.HandleFunc("/", handler_func)

	http.HandleFunc("/about/", handler_func_about)
	http.ListenAndServe(":8000", nil)
}

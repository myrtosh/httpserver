package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello Stranger!")
}

func main() {

	http.HandleFunc("/helloworld", hello)

	http.ListenAndServe(":8080", nil)
}

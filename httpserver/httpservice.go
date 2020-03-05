package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/fatih/camelcase"
)

func world(w http.ResponseWriter, r *http.Request) {
	value, _ := r.URL.Query()["name"]
	if len(value) > 0 {

		splitted := camelcase.Split(value[0])

		fmt.Fprintf(w, "Hello %v!", strings.Join(splitted, " "))
	} else if r.URL.Path[1:] == "helloworld" {
		fmt.Fprintln(w, "Hello Stranger!")
	}
}

func main() {
	http.HandleFunc("/", world)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// localhost:8080/helloworld?name=MyrtoSeisa
// localhost:8080/helloworld

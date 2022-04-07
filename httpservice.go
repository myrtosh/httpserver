package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Profile struct {
	GitHash          string
	NameOfTheProject string
}

var GitCommit string

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello Stranger!")
}

func version(w http.ResponseWriter, req *http.Request) {
	profile := Profile{GitCommit, "httpservice"}
	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content Type", "application/json")
	w.Write(js)
}
func main() {

	http.HandleFunc("/helloworld", hello)

	http.HandleFunc("/versionz", version)

	http.ListenAndServe(":8080", nil)
}

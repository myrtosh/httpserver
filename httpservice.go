package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Profile struct {
	GitHash          string
	NameOfTheProject string
}

var GitCommit string

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello Stranger!")
	CheckHTTPStatus()
}

func version(w http.ResponseWriter, req *http.Request) {
	profile := Profile{GitCommit, "httpservice"}
	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	CheckHTTPStatus()
}

func CheckHTTPStatus() {
	resp, _ := http.Get("http://localhost:8080")
	log.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		log.Println("HTTP Status is in the 2xx range")
	} else {
		log.Println("Argh! Broken")
	}
}

func main() {

	http.HandleFunc("/helloworld", hello)

	http.HandleFunc("/versionz", version)

	t := time.Now()
	fmt.Println(t.Format("2006-01-02-15:04:05"))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

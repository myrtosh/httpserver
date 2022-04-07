package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type Profile struct {
	GitHash          string
	NameOfTheProject string
}

var GitCommit string

func camelRegexp(str string) string {
	re := regexp.MustCompile(`([A-Z])`)
	//(^[^A-Z]*|[A-Z]*)([A-Z][^A-Z]+|$)
	str = re.ReplaceAllString(str, ` $1`)
	str = strings.Trim(str, " ")
	return str
}

func world(w http.ResponseWriter, r *http.Request) {
	value := r.URL.Query().Get("name")
	if len(value) > 0 {
		fmt.Fprintln(w, "Hello", camelRegexp(value))
		CheckHTTPStatus()
	} else if r.URL.Path[1:] == "helloworld" {
		fmt.Fprintf(w, "Hello Stranger!")
		CheckHTTPStatus()
	} else if r.URL.Path[1:] == "versionz" {
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

	http.HandleFunc("/", world)

	t := time.Now()
	fmt.Println(t.Format("2006-01-02-15:04:05"))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

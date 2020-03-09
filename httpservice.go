package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/fatih/camelcase"
)

type Profile struct {
	Git_hash            string
	Name_of_the_project string
}

func world(w http.ResponseWriter, r *http.Request) {
	value, _ := r.URL.Query()["name"]
	if len(value) > 0 {

		splitted := camelcase.Split(value[0])

		fmt.Fprintf(w, "Hello %v!", strings.Join(splitted, " "))
	} else if r.URL.Path[1:] == "helloworld" {
		fmt.Fprintln(w, "Hello Stranger!")
	} else if r.URL.Path[1:] == "versionz" {
		profile := Profile{"dd6492921d3fec85f9dcedfa58b20702b2e781a6", "httpservice"}

		js, err := json.Marshal(profile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}
func main() {
	http.HandleFunc("/", world)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const version string = "4.0.1"

var (
	tmpl = template.Must(template.ParseFiles("index.html"))
)

// PageData provides version details
type PageData struct {
	Version string
}

func main() {
	var port int

	flag.IntVar(&port, "p", 8080, "hosting port ")

	flag.Parse()

	fmt.Printf("Version %s\n", version)
	fmt.Printf("Port %d\n", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := &PageData{
			Version: version,
		}
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Version %s\n", version)
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

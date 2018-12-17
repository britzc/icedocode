package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

const version string = "6.0.1"

var (
	tmpl = template.Must(template.ParseFiles("index.html"))
)

// PageData provides version details
type PageData struct {
	Version     string
	NatsURL     string
	InfluxDBURL string
}

func main() {
	var port int

	flag.IntVar(&port, "p", 8080, "hosting port ")

	flag.Parse()

	fmt.Printf("Version %s\n", version)
	fmt.Printf("Port %d\n", port)

	natsURL := os.Getenv("NATS_HOST")
	influxDBURL := os.Getenv("INFLUXDB_HOST")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := &PageData{
			Version:     version,
			NatsURL:     natsURL,
			InfluxDBURL: influxDBURL,
		}
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Version %s\n", version)
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

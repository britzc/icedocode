package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	client "github.com/influxdata/influxdb/client/v2"
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

	fmt.Println("Connecting to InfluxDB")
	if err := writeData(influxDBURL); err != nil {
		fmt.Println(err)
	}

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

func writeData(influxDBHost string) (err error) {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     fmt.Sprintf("http://%s:8086", influxDBHost),
		Username: "jenkins",
		Password: "secretpassword",
	})
	if err != nil {
		return err
	}
	defer c.Close()

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "icedoapp",
		Precision: "s",
	})
	if err != nil {
		return err
	}

	// Create a point and add to batch
	tags := map[string]string{"cpu": "cpu-total"}
	fields := map[string]interface{}{
		"idle":   10.1,
		"system": 53.3,
		"user":   46.6,
	}

	pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
	if err != nil {
		return err
	}
	bp.AddPoint(pt)

	// Write the batch
	if err := c.Write(bp); err != nil {
		return err
	}

	// Close client resources
	if err := c.Close(); err != nil {
		return err
	}

	return nil
}

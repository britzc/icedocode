package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

const version string = "1.0.1"

func main() {
	var port int

	flag.IntVar(&port, "p", 8080, "hosting port ")

	flag.Parse()

	fmt.Printf("Version %s\n", version)
	fmt.Printf("Port %d\n", port)

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s\n", version)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

}

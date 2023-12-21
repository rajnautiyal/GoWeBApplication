package main

import (
	"fmt"
	"github/rajnautiyal/webCourse/cmd/pkg/handler"

	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)

	fmt.Printf("starting my application server %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}

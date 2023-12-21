package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprint(w, "hello world")

		if err != nil {
			fmt.Fprintf(os.Stderr, "Fprint: %v\n", err)
		}
		fmt.Print(n, " bytes written.\n")

	})
	_ = http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello Full Cycle</h1>"))
}

func main() {
	http.HandleFunc("/", Hello)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

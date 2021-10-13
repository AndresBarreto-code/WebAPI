package main

import (
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Running ok!")
}

func HandleAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from API!")
}

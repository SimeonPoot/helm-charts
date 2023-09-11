package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router instance
	router := mux.NewRouter()

	// Define a route handler for the root ("/") path
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	// Define another route handler for a custom path
	router.HandleFunc("/greet/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		fmt.Fprintf(w, "Hello, %s!", name)
	})

	// Start the HTTP server with the Gorilla Mux router
	fmt.Println("Listening on :8084...")
	err := http.ListenAndServe(":8084", router)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

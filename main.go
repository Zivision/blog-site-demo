package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got Request")
	io.WriteString(w, "This is the root page. Welcome!\n")
}

func main() {
	// Routes
	http.HandleFunc("/", getRoot)

	// Start server
	fmt.Println("Server Started")
	err := http.ListenAndServe(":8080", nil)

	// Error checking
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server closed")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got Request")
	io.WriteString(w, "This is the root page. Welcome!\n")
}

func main() {
	const PORT string = ":8080"
	// Routes
	http.HandleFunc("/", getRoot)

	// Start Server
	fmt.Println("Server running on localhost" + PORT);
	log.Fatal(http.ListenAndServe(PORT, nil))
}

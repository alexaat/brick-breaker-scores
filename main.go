package main

import (
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

	}
	http.HandleFunc("/score", scoreHandler)
	http.HandleFunc("/score/", scoreHandler)
	http.HandleFunc("/rank", rankHandler)
	http.HandleFunc("/", ping)
	//fmt.Println("Starting server on port :8080")
	http.ListenAndServe(":"+port, nil)
}

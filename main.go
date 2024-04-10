package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.HandleFunc("/score", scoreHandler)
	http.HandleFunc("/score/", scoreHandler)
	http.HandleFunc("/rank", rankHandler)
	http.HandleFunc("/", ping)
	fmt.Println("Starting server on port :" + port)
	http.ListenAndServe(":"+port, nil)
}

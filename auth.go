package main

import (
	"fmt"
	"net/http"
)

const (
	API_PATH = "/auth"
	API_PORT = 20000
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {

	fmt.Println("Authentication Server Started ...")
	http.HandleFunc(API_PATH, AuthHandler)
	serverAddress := fmt.Sprintf("localhost:%d", API_PORT)
	err := http.ListenAndServe(serverAddress, nil)
	if err == http.ErrServerClosed {
		fmt.Println("Server is Closed")
	} else {
		fmt.Println("Unknown Error , ERROR= ", err)
	}
}

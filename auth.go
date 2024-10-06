package main

import (
	"fmt"
	"net/http"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New request ...")
	w.WriteHeader(http.StatusOK)
}

func main() {

	fmt.Println("Authentication Server Started ...")
	http.HandleFunc("/auth", AuthHandler)
	err := http.ListenAndServe("localhost:2000", nil)
	if err == http.ErrServerClosed {
		fmt.Println("Server is Closed")
	}
}

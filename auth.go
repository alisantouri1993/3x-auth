package main

import (
	"fmt"
	"net/http"
	"slices"
)

const (
	API_PATH = "/auth"
	API_PORT = 20000
)

type UnameConnectionStore interface {
	CheckUname(username, ip string) int
}

type Uname struct {
	maxConnections   int
	connectedClients []string
}

type InMemoryUnameStore struct {
	store map[string]Uname
}

func (imStore *InMemoryUnameStore) CheckUname(username, ip string) int {
	var checkStatus int
	uname, ok := imStore.store[username]
	if ok {
		if len(uname.connectedClients) == uname.maxConnections {
			idx := slices.Index(uname.connectedClients, ip)
			if idx == -1 {
				checkStatus = http.StatusUnauthorized
			} else {
				checkStatus = http.StatusOK
			}
		} else {
			uname.connectedClients = append(uname.connectedClients, ip)
			imStore.store[username] = uname
			checkStatus = http.StatusOK
		}
	} else {
		checkStatus = http.StatusUnauthorized
	}
	return checkStatus
}

type AuthServer struct {
	store UnameConnectionStore
}

func (as *AuthServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var StatusCode int
	ip := r.Header.Get("X-Original-IP")
	username := r.Header.Get("X-Username")

	if ip != "" && username != "" {
		StatusCode = as.store.CheckUname(username, ip)
	} else {
		StatusCode = http.StatusBadRequest
	}
	w.WriteHeader(StatusCode)
}

func main() {

	fmt.Println("Authentication Server Started ...")
	store := InMemoryUnameStore{store: map[string]Uname{}}
	authserver := AuthServer{store: &store}
	http.HandleFunc(API_PATH, authserver.ServeHTTP)
	serverAddress := fmt.Sprintf("localhost:%d", API_PORT)
	err := http.ListenAndServe(serverAddress, nil)
	if err == http.ErrServerClosed {
		fmt.Println("Server is Closed")
	} else {
		fmt.Println("Unknown Error , ERROR= ", err)
	}
}

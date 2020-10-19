package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	serverList = []*server{
		launchServer("server-01", "http://127.0.0.1:3001"),
		launchServer("server-02", "http://127.0.0.1:3002"),
		launchServer("server-03", "http://127.0.0.1:3003"),
	}
	lastServedIndex = 0
)

func main() {
	http.HandleFunc("/", forwardRequest)
	go serverHealthCheck()
	log.Fatal(http.ListenAndServe(":6097", nil))
}

func forwardRequest(res http.ResponseWriter, req *http.Request) {
	server, err := getHealthyServer()
	if err != nil {
		http.Error(res, "Couldn't process request: "+err.Error(), http.StatusServiceUnavailable)
		return
	}
	server.ReverseProxy.ServeHTTP(res, req)
}

func getHealthyServer() (*server, error) {
	for i := 0; i < len(serverList); i++ {
		server := getServer()
		if server.Health {
			return server, nil
		}
	}
	return nil, fmt.Errorf("No healthy hosts")
}

func getServer() *server {
	nextIndex := (lastServedIndex + 1) % len(serverList)
	server := serverList[nextIndex]
	lastServedIndex = nextIndex
	return server
}

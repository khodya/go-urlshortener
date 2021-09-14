package main

import (
	"flag"
	"net/http"
	"os"
	"time"

	"github.com/khodya/go-urlshortener/internal/router"
)

var serverAddress string

func init() {
	flag.StringVar(&serverAddress, "a", parseServerAddress(), "Server address is not specified.")
}

func main() {
	flag.Parse()
	router := router.SetupRouter()
	s := &http.Server{
		Addr:           serverAddress,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func parseServerAddress() string {
	serverAddress, ok := os.LookupEnv("SERVER_ADDRESS")
	if !ok {
		serverAddress = ":8080"
	}
	return serverAddress
}

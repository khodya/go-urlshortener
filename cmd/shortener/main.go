package main

import (
	"net/http"
	"os"
	"time"

	"github.com/khodya/go-urlshortener/internal/router"
)

func main() {
	router := router.SetupRouter()
	serverAddress, ok := os.LookupEnv("SERVER_ADDRESS")
	if !ok {
		serverAddress = ":8080"
	}
	s := &http.Server{
		Addr:           serverAddress,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

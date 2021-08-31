package main

import (
	"github.com/khodya/go-urlshortener/internal/router"
)

func main() {
	router := router.SetupRouter()
	router.Run("localhost:8080")
}

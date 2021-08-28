package main

import (
	"github.com/gin-gonic/gin"
	"github.com/khodya/go-urlshortener/internal/app"
)

func main() {
	router := gin.Default()
	router.GET("/:id", app.Unfold)
	router.POST("/", app.Fold)

	router.Run("localhost:8080")
}

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/khodya/go-urlshortener/internal/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/:id", handlers.Unfold)
	r.POST("/", handlers.Fold)
	r.POST("/api/shorten", handlers.Shorten)
	return r
}

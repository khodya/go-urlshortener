package router

import (
	"github.com/gin-gonic/gin"
	"github.com/khodya/go-urlshortener/internal/app"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/:id", app.Unfold)
	r.POST("/", app.Fold)
	return r
}

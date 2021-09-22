package router

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	"github.com/khodya/go-urlshortener/internal/auth"
	"github.com/khodya/go-urlshortener/internal/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(Cookie())
	r.Use(gzip.Gzip(gzip.BestSpeed, gzip.WithDecompressFn(gzip.DefaultDecompressHandle)))
	r.GET("/:id", handlers.Unfold)
	r.POST("/", handlers.Fold)
	r.POST("/api/shorten", handlers.Shorten)
	return r
}

func Cookie() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("user")
		if err != nil || len(cookie) == 0 {
			c.SetCookie("user", auth.NewUserID(), 300, "", "", false, false)
		}
		c.Next()
	}
}

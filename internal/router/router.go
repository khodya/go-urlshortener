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
	r.GET("/user/urls", handlers.GetURLsByUser)
	return r
}

func Cookie() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("user")
		if err != nil || len(cookie) == 0 {
			userID := auth.NewUserID()
			c.SetCookie("user", userID, 300, "", "", false, false)
			c.Set("user", userID)
		}
		c.Next()
	}
}

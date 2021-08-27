package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func fold(c *gin.Context) {
	c.IndentedJSON(http.StatusAccepted, nil)
}

func unfold(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
}

func main() {
	router := gin.Default()
	router.GET("/:id", unfold)
	router.POST("/", fold)

	router.Run("localhost:8080")
}

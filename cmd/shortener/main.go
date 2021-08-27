package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func fold(c *gin.Context) {
	url, err := ioutil.ReadAll(c.Request.Body)
	if err != nil || len(url) == 0 {
		c.Status(400)
		return
	}
	c.IndentedJSON(http.StatusCreated, url)
}

func unfold(c *gin.Context) {
	c.Header("Location", "https://www.yandex.com")
	c.Status(http.StatusTemporaryRedirect)
}

func main() {
	router := gin.Default()
	router.GET("/:id", unfold)
	router.POST("/", fold)

	router.Run("localhost:8080")
}

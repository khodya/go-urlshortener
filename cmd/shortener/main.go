package main

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func fold(c *gin.Context) {
	url, err := ioutil.ReadAll(c.Request.Body)
	if err != nil || len(url) == 0 {
		c.Status(400)
		return
	}
	c.String(http.StatusCreated, "http://localhost:8080/asdfghj3453jhg3")
}

func unfold(c *gin.Context) {
	id := c.Param("id")
	url, err := url.Parse(id)
	if err != nil || url.Scheme == "" {
		c.IndentedJSON(http.StatusBadRequest, url)
	}
	c.Header("Location", "https://www.yandex.com")
	c.Status(http.StatusTemporaryRedirect)
}

func main() {
	router := gin.Default()
	router.GET("/:id", unfold)
	router.POST("/", fold)

	router.Run("localhost:8080")
}

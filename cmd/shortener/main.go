package main

import (
	"encoding/base64"
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
	c.String(http.StatusCreated, "http://localhost:8080/%s", base64.StdEncoding.EncodeToString(url))
}

func unfold(c *gin.Context) {
	id := c.Param("id")
	url, err := base64.StdEncoding.DecodeString(id)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad id parameter:%s", id)
	}
	c.Header("Location", string(url))
	c.Status(http.StatusTemporaryRedirect)
}

func main() {
	router := gin.Default()
	router.GET("/:id", unfold)
	router.POST("/", fold)

	router.Run("localhost:8080")
}

package handlers

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type shortenBody struct {
	Url string `json:"url"`
}

func Fold(c *gin.Context) {
	defer c.Request.Body.Close()
	url, err := io.ReadAll(c.Request.Body)
	if err != nil || len(url) == 0 {
		c.Status(400)
		return
	}
	c.String(http.StatusCreated, "http://localhost:8080/%s", base64.StdEncoding.EncodeToString(url))
}

func Unfold(c *gin.Context) {
	id := c.Param("id")
	url, err := base64.StdEncoding.DecodeString(id)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad id parameter:%s", id)
	}
	c.Header("Location", string(url))
	c.Status(http.StatusTemporaryRedirect)
}

func Shorten(c *gin.Context) {
	var body shortenBody
	if err := c.BindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, "Bad request body")
	}
	response := struct {
		Result string `json:"result"`
	}{
		Result: fmt.Sprintf("http://localhost:8080/%s", base64.StdEncoding.EncodeToString([]byte(body.Url))),
	}
	c.IndentedJSON(http.StatusCreated, response)
}

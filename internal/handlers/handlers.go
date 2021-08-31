package handlers

import (
	"encoding/base64"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type myRequestBody struct {
	URLText string `json:"url"`
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
	defer c.Request.Body.Close()
	body, err := io.ReadAll(c.Request.Body)
	if err != nil || len(body) == 0 {
		c.JSON(http.StatusBadRequest, struct{}{})
		return
	}
	var requestBody myRequestBody
	if err := json.Unmarshal(body, &requestBody); err != nil || len(requestBody.URLText) == 0 {
		c.JSON(http.StatusBadRequest, struct{}{})
		return
	}
	response := struct {
		Result string `json:"result"`
	}{
		Result: fmt.Sprintf("http://localhost:8080/%s", base64.StdEncoding.EncodeToString([]byte(requestBody.URLText))),
	}
	c.IndentedJSON(http.StatusCreated, response)
}

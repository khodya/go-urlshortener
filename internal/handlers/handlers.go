package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/khodya/go-urlshortener/internal/shortener"
)

type myRequestBody struct {
	URLText string `json:"url"`
}

const (
	BASE_URL_ENV_NAME = "BASE_URL"
	DEFAULT_BASE_URL  = "http://localhost:8080"
)

var baseUrl url.URL

func init() {
	baseUrl = parseBaseUrl()
}

func parseBaseUrl() url.URL {
	baseUrl, err := url.Parse(os.Getenv(BASE_URL_ENV_NAME))
	if err != nil || (*baseUrl == url.URL{}) {
		baseUrl, _ = url.Parse(DEFAULT_BASE_URL)
	}
	return *baseUrl
}

func Fold(c *gin.Context) {
	defer c.Request.Body.Close()
	url, err := io.ReadAll(c.Request.Body)
	if err != nil || len(url) == 0 {
		c.Status(400)
		return
	}
	resultUrl := baseUrl
	resultUrl.Path = shortener.Encode(url)
	c.String(http.StatusCreated, "%s", resultUrl.String())
	fmt.Printf("Base URL: %s\n", &baseUrl)
}

func Unfold(c *gin.Context) {
	id := c.Param("id")
	url, err := shortener.Decode(id)
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
	resultUrl := baseUrl
	resultUrl.Path = shortener.Encode([]byte(requestBody.URLText))
	c.IndentedJSON(http.StatusCreated, struct {
		Result string `json:"result"`
	}{
		Result: resultUrl.String(),
	})
}

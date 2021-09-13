package handlers

import (
	"encoding/json"
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
	BaseURLEnvName = "BASE_URL"
	DefaultBaseURL = "http://localhost:8080"
)

var baseURL url.URL

func init() {
	baseURL = parseBaseURL()
}

func parseBaseURL() url.URL {
	baseURL, err := url.Parse(os.Getenv(BaseURLEnvName))
	if err != nil || (*baseURL == url.URL{}) {
		baseURL, _ = url.Parse(DefaultBaseURL)
	}
	return *baseURL
}

func Fold(c *gin.Context) {
	defer c.Request.Body.Close()
	url, err := io.ReadAll(c.Request.Body)
	if err != nil || len(url) == 0 {
		c.Status(400)
		return
	}
	shortURL := composeURL(baseURL, shortener.Encode(url))
	c.String(http.StatusCreated, "%s", shortURL)
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
	shortURL := shortener.Encode([]byte(requestBody.URLText))
	c.IndentedJSON(http.StatusCreated, struct {
		Result string `json:"result"`
	}{
		Result: shortURL,
	})
}

func composeURL(baseURL url.URL, path string) string {
	url := baseURL
	url.Path = path
	return url.String()
}

package handlers

import (
	"encoding/json"
	"flag"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/khodya/go-urlshortener/internal/shortener"
	"github.com/khodya/go-urlshortener/internal/storage"
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
	baseURL = parseBaseURL(os.Getenv(BaseURLEnvName))
	flag.Func("b", "base url flag", func(flagValue string) error {
		baseURL = parseBaseURL(flagValue)
		return nil
	})
}

func parseBaseURL(urlToParse string) url.URL {
	baseURL, err := url.Parse(urlToParse)
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
	path := shortener.Encode(url)
	storage.Put(path, string(url))
	shortURL := composeURL(baseURL, path)
	userId, err := c.Cookie("user")
	if err == nil {
		storage.PutUser(userId, path)
	}
	c.String(http.StatusCreated, "%s", shortURL)
}

func Unfold(c *gin.Context) {
	id := c.Param("id")
	decodedURL, ok := storage.Get(id)
	if !ok {
		url, err := shortener.Decode(id)
		if err != nil {
			c.String(http.StatusBadRequest, "Bad id parameter:%s", id)
		}
		decodedURL = string(url)
	}

	c.Header("Location", decodedURL)
	c.Status(http.StatusTemporaryRedirect)
}

func Shorten(c *gin.Context) {
	defer c.Request.Body.Close()
	body, err := io.ReadAll(c.Request.Body)
	if err != nil || len(body) == 0 {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	var requestBody myRequestBody
	if err := json.Unmarshal(body, &requestBody); err != nil || len(requestBody.URLText) == 0 {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	path := shortener.Encode([]byte(requestBody.URLText))
	storage.Put(path, requestBody.URLText)
	shortURL := composeURL(baseURL, path)
	userId, err := c.Cookie("user")
	if err == nil {
		storage.PutUser(userId, path)
	}
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

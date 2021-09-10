package shortener

import (
	"encoding/base64"
	"net/url"
	"os"
)

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

func Encode(v []byte) string {
	url := baseURL
	url.Path = base64.StdEncoding.EncodeToString(v)
	return url.String()
}

func Decode(v string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(v)

}

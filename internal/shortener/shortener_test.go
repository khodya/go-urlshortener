package shortener

import (
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	type want struct {
		baseURL     string
		urlToEncode string
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "Encode. Happy",
			want: want{
				baseURL:     "http://localhost:1234",
				urlToEncode: "https://yandex.com",
			},
		},
		{
			name: "Encode. Happy with slash",
			want: want{
				baseURL:     "http://localhost:1234/",
				urlToEncode: "https://yandex.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv(BaseURLEnvName, tt.want.baseURL)

			actualURLRaw := Encode([]byte(tt.want.urlToEncode))
			assert.NotEmpty(t, actualURLRaw)

			actualURL, err := url.Parse(actualURLRaw)
			assert.Nil(t, err)
			assert.NotEmpty(t, actualURL)

			os.Unsetenv(BaseURLEnvName)
		})
	}
}

func TestDecode(t *testing.T) {
	myURL := "https://yandex.com"
	encodedURL := Encode([]byte(myURL))

	parsedEncodedURL, _ := url.Parse(encodedURL)
	pathParameter := strings.Replace(parsedEncodedURL.Path, "/", "", 1)
	decodedURL, err := Decode(pathParameter)

	assert.Nil(t, err)
	assert.NotEmpty(t, decodedURL)
	assert.Equal(t, myURL, string(decodedURL))
}

func TestParseBaseURL(t *testing.T) {
	type want struct {
		envVar      string
		expectedURL string
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "parseBaseURL. Happy",
			want: want{
				envVar:      "http://localhost:1234",
				expectedURL: "http://localhost:1234",
			},
		},
		{
			name: "parseBaseURL. Default",
			want: want{
				envVar:      ":::sdfdsf:fsdfdf:fdsfsd:",
				expectedURL: DefaultBaseURL,
			},
		},
		{
			name: "parseBaseURL. Empty",
			want: want{
				envVar:      "",
				expectedURL: DefaultBaseURL,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv(BaseURLEnvName, tt.want.envVar)

			parsedURL := parseBaseURL()
			assert.Equal(t, tt.want.expectedURL, parsedURL.String())

			os.Unsetenv(BaseURLEnvName)
		})
	}
}

package handlers

import (
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestComposeURL(t *testing.T) {
	type want struct {
		url         url.URL
		path        string
		expectedURL string
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "composeURL. Happy",
			want: want{
				url:         url.URL{Scheme: "http", Host: "localhost"},
				path:        "123",
				expectedURL: "http://localhost/123",
			},
		},
		{
			name: "composeURL. happy with slash",
			want: want{
				url:         url.URL{Scheme: "http", Host: "localhost"},
				path:        "/123",
				expectedURL: "http://localhost/123",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			composedURL := composeURL(tt.want.url, tt.want.path)
			assert.Equal(t, tt.want.expectedURL, composedURL)
		})
	}
}

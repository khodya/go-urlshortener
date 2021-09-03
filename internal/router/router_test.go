package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFold(t *testing.T) {
	router := SetupRouter()

	type want struct {
		code int
		url  string
		body string
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "happy. fold",
			want: want{
				code: 201,
				url:  "/",
				body: "https://www.yandex.com",
			},
		},
		{
			name: "unhappy. fold 404",
			want: want{
				code: 404,
				url:  "/name",
				body: "https://www.yandex.com",
			},
		},
		{
			name: "unhappy. fold 400",
			want: want{
				code: 400,
				url:  "/",
				body: "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", tt.want.url, strings.NewReader(tt.want.body))
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.want.code, w.Code)
		})
	}
}

func TestUnfold(t *testing.T) {
	router := SetupRouter()

	type want struct {
		code int
		url  string
	}

	tests := []struct {
		name string
		want want
	}{
		{
			name: "happy. unfold",
			want: want{
				code: 307,
				url:  "/aHR0cHM6Ly93d3cueWFuZGV4LmNvbQ==",
			},
		},
		{
			name: "unhappy. unfold",
			want: want{
				code: 400,
				url:  "/aHR0cHM6Ly93d3cueWFuZGLmNvbQ==",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", tt.want.url, nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.want.code, w.Code)
		})
	}
}

func TestShorten(t *testing.T) {
	router := SetupRouter()

	type want struct {
		code int
		body string
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "Shorten. Happy",
			want: want{
				code: 201,
				body: "{\"url\": \"https://yandex.ru\"}",
			},
		},
		{
			name: "Shorten. Empty JSON",
			want: want{
				code: 400,
				body: "{}",
			},
		},
		{
			name: "Shorten. Empty request body",
			want: want{
				code: 400,
				body: "",
			},
		},
		{
			name: "Shorten. Invalid JSON",
			want: want{
				code: 400,
				body: "{\"url\": https://yandex.ru\"}",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/api/shorten", strings.NewReader(tt.want.body))
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.want.code, w.Code)
			assert.Equal(t, "application/json; charset=utf-8", w.Result().Header.Get("Content-Type"))

			// assert response is a valid JSON
			var js json.RawMessage
			assert.Nil(t, json.Unmarshal(w.Body.Bytes(), &js))

		})
	}
}

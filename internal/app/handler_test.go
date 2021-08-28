package app

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/:id", Unfold)
	r.POST("/", Fold)
	return r
}

func TestFold(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", strings.NewReader(""))
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestUnfold(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/aHR0cHM6Ly93d3cueWFuZGV4LmNvbQ==", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 307, w.Code)
}

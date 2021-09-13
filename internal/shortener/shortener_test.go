package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	myURL := "https://yandex.com"
	encodedURL := Encode([]byte(myURL))
	decodedURL, err := Decode(encodedURL)

	assert.Nil(t, err)
	assert.NotEmpty(t, decodedURL)
	assert.Equal(t, myURL, string(decodedURL))
}

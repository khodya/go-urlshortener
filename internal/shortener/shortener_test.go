package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecode(t *testing.T) {
	myURL := "https://yandex.com"
	encodedURL := Encode([]byte(myURL))
	decodedURL, err := Decode(encodedURL)

	require.NoError(t, err)
	assert.NotEmpty(t, decodedURL)
	assert.Equal(t, myURL, string(decodedURL))
}

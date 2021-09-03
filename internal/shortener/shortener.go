package shortener

import "encoding/base64"

func Encode(v []byte) string {
	return base64.StdEncoding.EncodeToString(v)
}

func Decode(v string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(v)

}

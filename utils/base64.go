package utils

import "encoding/base64"

func EncodeBase64(data string) string {
	encodedString := base64.StdEncoding.EncodeToString([]byte(data))
	return encodedString
}

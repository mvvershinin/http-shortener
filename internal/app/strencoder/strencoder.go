package strencoder

import (
	"encoding/base64"
)

func Base64Encode(str string) string {
	enc := base64.StdEncoding.EncodeToString([]byte(str))

	return enc
}

func Base64Decode(str string) (string, error) {
	dec, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return str, err
	}

	return string(dec), nil
}

package strencoder

import "encoding/base32"

func EncodeStr(stringToEncode string) string {
	str := base32.StdEncoding.EncodeToString([]byte(stringToEncode))

	return str
}

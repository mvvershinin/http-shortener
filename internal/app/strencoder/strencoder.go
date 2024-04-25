package strencoder

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func EncodeStr(stringToEncode string) string {
	hash := sha256.New()
	hash.Write([]byte(stringToEncode))
	hashSum := hash.Sum(nil)
	hashHex := hex.EncodeToString(hashSum)

	return strings.ToLower(hashHex)
}

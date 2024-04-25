package strencoder

import (
	"encoding/base64"
	"fmt"
)

func Base64Encode(str string) string {
	enc := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Printf("%s\n", enc)
	return enc
	//todo return err
}

func Base64Decode(str string) (string, error) {
	fmt.Println(str)
	dec, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return str, err
	}
	fmt.Printf("%s\n", string(dec))
	return string(dec), nil
	//todo return err
}

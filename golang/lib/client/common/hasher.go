package common

import (
	"crypto/sha1"
	"fmt"
)

func HashString(value string) string {
	hash := sha1.New()

	hash.Write([]byte(value))

	hashedByteSlice := hash.Sum(nil)

	hexValue := fmt.Sprintf("%x", hashedByteSlice)

	return hexValue
}

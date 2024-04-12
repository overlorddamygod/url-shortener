package utils

import (
	"crypto/sha256"

	"github.com/jxskiss/base62"
)

func GetHash(str string) []byte {
	hash := sha256.New()
	hash.Write([]byte(str))
	return hash.Sum(nil)
}

func EncodeBase62(src []byte, chars uint) string {
	return base62.EncodeToString(src)[:chars]
}

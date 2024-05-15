package hash

import (
	"crypto/sha1"
	"fmt"
)

var (
	salt = "qwerty0oi123sdjnxci0jk0oqi0jtgji9df123ko"
)

func GenerateSha1Hash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

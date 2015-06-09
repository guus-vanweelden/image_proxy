package filter

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"strings"
)

func simpleGenerateHash(str string) string {
	h := sha1.New()
	io.WriteString(h, strings.ToLower(str))
	hash := fmt.Sprintf("%x", h.Sum(nil))[5:25]
	return hash
}

func calcCacheKey(link string, method string, optionals string) string {
	key := strings.Join([]string{link, method, optionals}, "_")
	key = strings.ToLower(key)
	log.Printf("Key: %s", key)

	return simpleGenerateHash(key)
}

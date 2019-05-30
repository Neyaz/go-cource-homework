package shortener

import (
  "crypto/sha1"
	"encoding/hex"
	"regexp"
)

// Shortener has methods for short url and resolve url
type Shortener interface {
  Shorten(string) string
  Resolve(string) string
}

var urlAccociation  = map[string]string{}
var re = regexp.MustCompile(`(?:https?:\/\/)?(?:[a-zA-Z0-9.-]+?\.(?:[a-zA-Z])+)\/?`)

// Shorten takes url and return short url
func Shorten(rawURL string) string {
	var hash string;
	var shortURL string;
	matchArr := re.FindStringSubmatch(rawURL)
	path := re.ReplaceAllString(rawURL, "")
	if path == "" {
		shortURL = rawURL
	} else {
		hash = hasher(rawURL)
		shortURL = matchArr[0] + hash
		urlAccociation[hash] = rawURL
	}
	return shortURL
}

// Resolve takes short url and return orignal url
func Resolve(rawURL string) string {
	path := re.ReplaceAllString(rawURL, "")
	if path == "" {
		return rawURL
	}
	return urlAccociation[path]
}

func hasher(text string) string {
  	algorithm := sha1.New()
    algorithm.Write([]byte(text))
		return hex.EncodeToString(algorithm.Sum(nil))[0:4]
}
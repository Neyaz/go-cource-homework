package shortener

import (
  "crypto/sha1"
  "encoding/hex"
  "regexp"
)

// Shortener has methods for shorten url and resolve url
type Shortener interface {
  Shorten(string) string
  Resolve(string) string
}

// MyShortener has storage map
type MyShortener struct { 
  storage map[string]string
}

var re = regexp.MustCompile(`(?:https?:\/\/)?(?:[a-zA-Z0-9.-]+?\.(?:[a-zA-Z])+)\/?`)

// Shorten takes url and return short url
func (s MyShortener) Shorten(rawURL string) string {
  var hash string;
  var shortURL string;
  matchArr := re.FindStringSubmatch(rawURL)
  path := re.ReplaceAllString(rawURL, "")
  if path == "" {
    shortURL = rawURL
  } else {
    hash = s.hasher(rawURL)
    shortURL = matchArr[0] + hash
    s.storage[hash] = rawURL
  }
  return shortURL
}

// Resolve takes short url and return orignal url
func (s MyShortener) Resolve(rawURL string) string {
  path := re.ReplaceAllString(rawURL, "")
  if path == "" {
    return rawURL
  }
  return s.storage[path]
}

func (s MyShortener) hasher(text string) string {
  algorithm := sha1.New()
  algorithm.Write([]byte(text))
  return hex.EncodeToString(algorithm.Sum(nil))[0:4]
}
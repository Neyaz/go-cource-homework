package shortener

import (
  "testing"
)

var testCases = []struct {
  url        string
  shortenURL string
}{
  {
    url: "https://otus.ru/some_path?q=23&t=d",
    shortenURL: "https://otus.ru/f936",
  },
  {
    url: "otus.ru/some_path?q=23&t=d",
    shortenURL: "otus.ru/bd57",
  },
  {
    url: "https://otus.ru/",
    shortenURL: "https://otus.ru/",
  },
  {
    url: "otus.ru",
    shortenURL: "otus.ru",
  },
}

func TestShorten(t *testing.T) {
	s := MyShortener{storage: make(map[string]string)}
  for _, c := range testCases {
    actual := s.Shorten(c.url)
    if actual != c.shortenURL {
      t.Fatalf("FAIL: URL %q, expected %q, got %q", c.url, c.shortenURL, actual)
    }

    t.Logf("PASS: URL %q", c.url)
  }
}

func TestResolve(t *testing.T) {
	s := MyShortener{storage: make(map[string]string)}
  for _, c := range testCases {
    s.Shorten(c.url)
  }

  for _, c := range testCases {
    actual := s.Resolve(c.shortenURL)
    if actual != c.url {
      t.Fatalf("FAIL: URL %q, expected %q, got %q", c.shortenURL, c.url, actual)
    }

    t.Logf("PASS: URL %q", c.shortenURL)
  }
}
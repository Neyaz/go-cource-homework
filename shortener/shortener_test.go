package shortener

import (
	"testing"
)

var testCases = []struct {
	url       	string
	shortenUrl string
}{
	{
		url: "https://otus.ru/some_path?q=23&t=d",
		shortenUrl: "https://otus.ru/f936",
	},
	{
		url: "otus.ru/some_path?q=23&t=d",
		shortenUrl: "otus.ru/bd57",
	},
	{
		url: "https://otus.ru/",
		shortenUrl: "https://otus.ru/",
	},
	{
		url: "otus.ru",
		shortenUrl: "otus.ru",
	},
}

func TestShorten(t *testing.T) {
	for _, c := range testCases {
		actual := Shorten(c.url)
		if actual != c.shortenUrl {
			t.Fatalf("FAIL: URL %q, expected %q, got %q", c.url, c.shortenUrl, actual)
		}

		t.Logf("PASS: URL %q", c.url)
	}
}

func TestResolve(t *testing.T) {
	for _, c := range testCases {
		Shorten(c.url)
	}

	for _, c := range testCases {
		actual := Resolve(c.shortenUrl)
		if actual != c.url {
			t.Fatalf("FAIL: URL %q, expected %q, got %q", c.shortenUrl, c.url, actual)
		}

		t.Logf("PASS: URL %q", c.shortenUrl)
	}
}
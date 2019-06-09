package unpack

import (
  "testing"
)

var testCases = []struct {
  input  string
  result string
}{
  {
    input: "a4bc2d5e",
    result: "aaaabccddddde",
  },
  {
    input: "abcd",
    result: "abcd",
  },
  {
    input: "45",
    result: "",
  },
}

func TestUnpack(t *testing.T) {
  for _, c := range testCases {
		actual := Unpack(c.input)
    if actual != c.result {
      t.Fatalf("FAIL: Input %q, expected %q, got %q", c.input, c.result, actual)
    }

    t.Logf("PASS: URL %q", c.input)
  }
}
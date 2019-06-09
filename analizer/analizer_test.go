package analizer

import (
	"reflect"
	"testing"
)

var analizeTestCases = []struct {
	text     string
	list     WordsList
}{
	{
		text:     "one one one two",
		list:     []string{"one", "two"},
	},
}

func TestAnalize(t *testing.T) {
	for _, tt := range analizeTestCases {
		in := WordsList{}
		got := in.Analize(tt.text)
		if !reflect.DeepEqual(WordsList(tt.list), got) {
			t.Fatalf("FAIL: %s -- expected: %v, actual: %v", tt.text, tt.list, got)
		} else {
			t.Logf("PASS: %s", tt.text)
		}

	}
}
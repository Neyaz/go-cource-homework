package maximum

import (
  "testing"
)

func compare(a interface{}, b interface{}) interface{} {
  intA := a.(int)
  intB := b.(int)
  if intA < intB { 
    return intB
  }
  return intA
}

var maximumTestCases = []struct {
  fn       compareFunc
  list     []interface{}
  want     interface{}
}{
  {
    fn:       compare,
    list:     []interface{}{1,2,3,5,4},
    want:     5,
  },
}

func TestMaximumMethod(t *testing.T) {
  for _, tt := range maximumTestCases {
    in := AnyList(tt.list)
    got := in.Maximum(tt.fn)
    if tt.want != got {
      t.Fatalf("FAIL: expected: %v, actual: %v", tt.want, got)
    } else {
      t.Logf("PASS: %s", got)
    }

  }
}
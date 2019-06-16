package maximum

import (
  "log"
)

// AnyList list of some elements
type AnyList []interface{}

type compareFunc func(interface{}, interface{}) interface{}

// Maximum find maximum in list
func (l AnyList) Maximum(comp compareFunc) interface{} {
  var max interface{}

  if len(l) == 0 {
    log.Fatal("Empty slice")
  }
  
  max = l[0]
  for _, value := range l {
    max = comp(max, value)
  }

  return max;
}
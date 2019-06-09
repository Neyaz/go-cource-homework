package analizer

import (
	"strings"
	"sort"
)

// WordsList contains most frequent words
type WordsList []string

// Analize recieves text and return most frequent words
func (wl WordsList) Analize(text string) WordsList {
	words := strings.Fields(text)
	counts := make(map[string]int, len(words))
	freq := make(map[int][]string, len(words))
	for _, word := range words {
		counts[word]++
	}
	for k, v := range counts {
      freq[v] = append(freq[v], k)
	}
	keys := make([]int, 0, len(counts))
	for k := range freq {
        keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	for _, v := range keys {
    for _, s := range freq[v] {
			if len(wl) <= 10 {
				wl = append(wl, s)
			}
		}
	}
	return wl
}

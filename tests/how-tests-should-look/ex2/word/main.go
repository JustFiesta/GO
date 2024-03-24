package word

import "strings"

// UseCount returns number of times each words is used in a string
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns number of words in given string
func Count(s string) int {
	count := len(strings.Fields(s))
	return count
}

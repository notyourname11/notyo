package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, v := range strings.Fields(s) {
		_, ok := m[v]
		if !ok {
			m[v] = 0
		}
		m[v]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}

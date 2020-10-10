package utils

import (
	"fmt"
	"strings"
	"testing"
)

func TestGenRandString(t *testing.T) {
	m := make(map[string]int)

	count := 10000
	for i := 0; i < count; i++ {
		s := GenRandString(6)
		s = strings.ToLower(s)
		m[s]++
	}

	for k, v := range m {
		s := fmt.Sprintf("k: %s, v: %d", k, v)
		if v != 1 {
			fmt.Printf("%s -------------------\n", s)
		} else {
			fmt.Printf("%s\n", s)
		}
	}
}

func TestUUID(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s\n", UUID())
	}
}

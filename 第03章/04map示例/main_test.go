package main

import (
	"testing"
)

func TestLengthOfNonRepeatingSubString(t *testing.T) {
	tests := []struct {
		s string
		n int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"b", 1},
		{"abcdefg", 7},
		{"锄禾日当午", 5},
		{"一二三二一", 3},
	}

	for _, tt := range tests {
		if actual := lengthOfNonRepeatingSubString(tt.s); actual != tt.n {
			t.Errorf("%q got %d, expected %d\n", tt.s, actual, tt.n)
		}
	}
}

func BenchmarkLengthOfNonRepeatingSubString(b *testing.B) {
	s, n := "一二三二一", 3
	for i := 0; i < b.N; i++ {
		if actual := lengthOfNonRepeatingSubString(s); actual != n {
			b.Errorf("%q got %d, expected %d\n", s, actual, n)
		}
	}
}

package main

import (
	"testing"
)

func TestDecodeString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Example 1 - Single Layer",
			input:    "3[a]2[bc]",
			expected: "aaabcbc",
		},
		{
			name:     "Example 2 - Nested",
			input:    "3[a2[c]]",
			expected: "accaccacc",
		},
		{
			name:     "Example 3 - Mixed and Sequential",
			input:    "2[abc]3[cd]ef",
			expected: "abcabccdcdcdef",
		},
		{
			name:     "No Digits",
			input:    "leetcode",
			expected: "leetcode",
		},
		{
			name:     "Deeply Nested",
			input:    "2[2[2[b]]]",
			expected: "bbbbbbbb",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeString(tt.input); got != tt.expected {
				t.Errorf("decodeString(%q) = %v, expected %v", tt.input, got, tt.expected)
			}
		})
	}
}

package main

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "Example 1: Simple valid pair",
			s:    "()",
			want: true,
		},
		{
			name: "Example 2: Multiple consecutive valid pairs",
			s:    "()[]{}",
			want: true,
		},
		{
			name: "Example 3: Mismatched pair",
			s:    "(]",
			want: false,
		},
		{
			name: "Example 4: Nested valid pairs",
			s:    "([])",
			want: true,
		},
		{
			name: "Example 5: Interleaved invalid pairs",
			s:    "([)]",
			want: false,
		},
		{
			name: "Edge Case: Odd length string",
			s:    "()[",
			want: false,
		},
		{
			name: "Edge Case: Only left brackets",
			s:    "(((",
			want: false,
		},
		{
			name: "Edge Case: Only right brackets",
			s:    "]]]",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.s); got != tt.want {
				t.Errorf("isValid(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

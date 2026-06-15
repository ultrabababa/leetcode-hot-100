package main

import (
	"testing"
)

func TestMinStack(t *testing.T) {
	tests := []struct {
		name     string
		commands []string
		values   [][]int
		expected []interface{} // 使用 interface{} 兼容 int 和 nil
	}{
		{
			name:     "Example 1",
			commands: []string{"MinStack", "push", "push", "push", "getMin", "pop", "top", "getMin"},
			values:   [][]int{{}, {-2}, {0}, {-3}, {}, {}, {}, {}},
			expected: []interface{}{nil, nil, nil, nil, -3, nil, 0, -2},
		},
		{
			name:     "All Same Values",
			commands: []string{"MinStack", "push", "push", "getMin", "pop", "getMin"},
			values:   [][]int{{}, {1}, {1}, {}, {}, {}},
			expected: []interface{}{nil, nil, nil, 1, nil, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var obj MinStack
			for i, cmd := range tt.commands {
				switch cmd {
				case "MinStack":
					obj = Constructor()
				case "push":
					obj.Push(tt.values[i][0])
				case "pop":
					obj.Pop()
				case "top":
					if got := obj.Top(); got != tt.expected[i] {
						t.Errorf("step %d (top): expected %v, got %v", i, tt.expected[i], got)
					}
				case "getMin":
					if got := obj.GetMin(); got != tt.expected[i] {
						t.Errorf("step %d (getMin): expected %v, got %v", i, tt.expected[i], got)
					}
				}
			}
		})
	}
}

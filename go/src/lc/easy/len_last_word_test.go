package easy

import "testing"

func TestLenLastWord(t *testing.T) {
	tests := []struct {
		str      string
		expected int
	}{
		{"a", 1},
		{"a ", 1},
		{"Hello World", 5},
		{"day", 3},
	}

	for _, tt := range tests {
		if got := lengthOfLastWord(tt.str); got != tt.expected {
			t.Errorf("lengthOfLastWord(%v), got=%v, expected=%v", tt.str, got, tt.expected)
		}

	}
}

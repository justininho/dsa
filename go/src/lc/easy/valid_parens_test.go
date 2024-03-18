package easy

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"{[]", false},
		{"{[}]", false},
	}
	for _, tt := range tests {
		if got := isValid(tt.s); got != tt.want {
			t.Errorf("isValid(%v) = %v; want = %v", tt.s, got, tt.want)
		}
	}
}

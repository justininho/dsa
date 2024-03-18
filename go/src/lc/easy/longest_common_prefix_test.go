package easy

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		strs []string
		want string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		//{[]string{"", "b"}, ""},
	}

	for _, tt := range tests {
		if got := longestCommonPrefix(tt.strs); got != tt.want {
			t.Errorf("longestCommonPrefix(%v) = %v, want %v", tt.strs, got, tt.want)
		}
	}
}

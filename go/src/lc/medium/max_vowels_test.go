package medium

import "testing"

func TestMaxVowels(t *testing.T) {
	tests := []struct {
		s    string
		k    int
		want int
	}{
		//{"abciiidef", 3, 3},
		//{"aeiou", 2, 2},
		//{"leetcode", 3, 2},
		//{"rhythms", 4, 0},
		//{"tryhard", 4, 1},
		{"ibpbhixfiouhdljnjfflpapptrxgcomvnb", 33, 7},
	}
	for _, tt := range tests {
		if got := maxVowels(tt.s, tt.k); got != tt.want {
			t.Errorf("maxVowels(%v, %v) = %v, want %v", tt.s, tt.k, got, tt.want)
		}
	}
}

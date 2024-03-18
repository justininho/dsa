package easy

import "testing"

func TestTribonacci(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{4, 4},
		{25, 1389537},
	}
	for _, tt := range tests {
		if got := tribonacci(tt.n); got != tt.want {
			t.Errorf("tribonacci(%v) = %v, want %v", tt.n, got, tt.want)
		}
	}
}

package lc

import (
	"slices"
	"testing"
)

func TestGenParens(t *testing.T) {
	tests := []struct {
		n   int
		exp []string
	}{
		{0, []string{""}},
		{1, []string{"()"}},
		{2, []string{"()()", "(())"}},
		{3, []string{"()()()", "()(())", "(())()", "(()())", "((()))"}},
	}
	for _, tt := range tests {
		if res := generateParenthesis(tt.n); !slices.Equal(res, tt.exp) {
			t.Fatalf("expected %v, got %v", tt.exp, res)
		}
	}
}

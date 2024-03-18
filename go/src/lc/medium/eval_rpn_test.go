package medium

import "testing"

func TestEvalRpn(t *testing.T) {
	inputs := []struct {
		tokens   []string
		expected int
	}{
		{[]string{"2", "1", "+", "3", "*"}, 9},
		{[]string{"4", "13", "5", "/", "+"}, 6},
		{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
		{[]string{"3", "-4", "+"}, -1},
	}
	for _, input := range inputs {
		if got := evalRPN(input.tokens); got != input.expected {
			t.Errorf("evalRPN(%v) = %v; want = %v", input.tokens, got, input.expected)
		}
	}
}

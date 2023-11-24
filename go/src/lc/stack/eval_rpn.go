package lc

import (
	"dsa/src/datastructures/stack"
	"strconv"
)

/*
(150 | Medium) evalRPN
*/
func evalRPN(tokens []string) int {
	s := stack.New[int]()
	for _, str := range tokens {
		if n, err := strconv.Atoi(str); err == nil {
			s.Push(n)
		} else {
			op2, _ := s.Pop()
			op1, _ := s.Pop()
			result := calc(op1, op2, str)
			s.Push(result)
		}
	}
	res, _ := s.Pop()
	return res
}

func calc(operand1 int, operand2 int, operation string) int {
	switch operation[0] {
	case '+':
		return operand1 + operand2
	case '-':
		return operand1 - operand2
	case '*':
		return operand1 * operand2
	case '/':
		return operand1 / operand2
	default:
		return 0
	}
}

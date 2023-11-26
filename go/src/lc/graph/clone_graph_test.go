package lc

import "testing"

func TestCloneGraph(t *testing.T) {
	var tests = []struct {
		node *Node
		want *Node
	}{
		{nil, nil},
		{&Node{Val: 1, Neighbors: []*Node{}}, &Node{Val: 1, Neighbors: []*Node{}}},
		{&Node{Val: 1, Neighbors: []*Node{&Node{Val: 2}}}, &Node{Val: 1, Neighbors: []*Node{&Node{Val: 2}}}},
		{&Node{Val: 1, Neighbors: []*Node{&Node{Val: 2}, &Node{Val: 3}}}, &Node{Val: 1, Neighbors: []*Node{&Node{Val: 2}, &Node{Val: 3}}}},
		{&Node{Val: 1, Neighbors: []*Node{&Node{Val: 2, Neighbors: []*Node{&Node{Val: 3}}}}}, &Node{Val: 1, Neighbors: []*Node{&Node{Val: 2, Neighbors: []*Node{&Node{Val: 3}}}}}},
	}

	for _, tt := range tests {
		got := cloneGraph(tt.node)
		if !equalGraph(got, tt.want) {
			t.Errorf("cloneGraph(%v) return %v, want %v", tt.node, got, tt.want)
		}
	}
}

func equalGraph(got, want *Node) bool {
	if got == nil && want == nil {
		return true
	}
	if got == nil || want == nil {
		return false
	}
	if got.Val != want.Val {
		return false
	}
	if len(got.Neighbors) != len(want.Neighbors) {
		return false
	}
	for i := 0; i < len(got.Neighbors); i++ {
		if !equalGraph(got.Neighbors[i], want.Neighbors[i]) {
			return false
		}
	}
	return true
}

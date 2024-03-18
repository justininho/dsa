package medium

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

type Node struct {
	Val       int
	Neighbors []*Node
}

/*
 * (133 | Medium) cloneGraph
 */
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := make(map[*Node]*Node)

	q := []*Node{node}
	visited[node] = &Node{Val: node.Val}
	for len(q) > 0 {
		n := q[0]
		for _, neighbor := range n.Neighbors {
			if _, has := visited[neighbor]; !has {
				visited[neighbor] = &Node{Val: neighbor.Val}
				q = append(q, neighbor)
			}
			visited[n].Neighbors = append(visited[n].Neighbors, visited[neighbor])
		}
		q = q[1:]
	}
	return visited[node]
}

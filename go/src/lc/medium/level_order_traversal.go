package medium

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	var q = []*TreeNode{root}
	for len(q) > 0 {
		var level []int
		for i := len(q); i > 0; i-- {
			node := q[0]
			if node != nil {
				level = append(level, node.Val)
				if node.Left != nil {
					q = append(q, node.Left)
				}
				if node.Right != nil {
					q = append(q, node.Right)
				}
			}
			q = q[1:]
		}
		if len(level) > 0 {
			result = append(result, level)
		}
	}
	return result
}

/*
Queue Implementation - as fast as the above solution
func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	var q = queue.New[*TreeNode]()
	q.Enqueue(root)
	for q.Len() > 0 {
		var level []int
		for i := q.Len(); i > 0; i-- {
			node, _ := q.Dequeue()
			if node != nil {
				level = append(level, node.Val)
				if node.Left != nil {
					q.Enqueue(node.Left)
				}
				if node.Right != nil {
					q.Enqueue(node.Right)
				}
			}
		}
		if len(level) > 0 {
			result = append(result, level)
		}
	}
	return result
}
*/

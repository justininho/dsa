package medium

/*
 * (210 | Medium) findOrder
 */
func findOrder(numCourses int, prerequisites [][]int) []int {
	requirements := make([]int, numCourses)
	dependencies := make([][]int, numCourses)

	for _, pr := range prerequisites {
		c, r := pr[0], pr[1]
		dependencies[r] = append(dependencies[r], c)
		requirements[c]++
	}

	var q []int
	for course, reqs := range requirements {
		if reqs == 0 {
			q = append(q, course)
		}
	}

	var result []int
	for len(q) > 0 {
		node := q[0]
		result = append(result, node)
		for _, neighbor := range dependencies[node] {
			requirements[neighbor]--
			if requirements[neighbor] == 0 {
				q = append(q, neighbor)
			}
		}
		q = q[1:]
	}
	if len(result) != numCourses {
		return []int{}
	}
	return result
}

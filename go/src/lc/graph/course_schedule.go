package lc

/*
 * (207 | Medium) canFinish
 */
func canFinish(numCourses int, prerequisites [][]int) bool {
	requirements := make([]int, numCourses)
	dependencies := make([][]int, numCourses)

	for _, pr := range prerequisites {
		c, r := pr[0], pr[1]
		dependencies[r] = append(dependencies[r], c)
		requirements[c]++
	}

	var queue []int
	for course, degree := range requirements {
		if degree == 0 {
			queue = append(queue, course)
		}
	}

	nodesVisited := 0
	for len(queue) > 0 {
		node := queue[0]
		nodesVisited++
		for _, neighbor := range dependencies[node] {
			requirements[neighbor]--
			if requirements[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
		queue = queue[1:]
	}
	return nodesVisited == numCourses
}

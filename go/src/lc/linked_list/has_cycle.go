package lc

/**
 * (141 | Easy) hasCycle
 */
func hasCycle(head *ListNode) bool {
	seen := make(map[*ListNode]bool)
	for head != nil {
		if _, has := seen[head]; has {
			return true
		}
		seen[head] = true
		head = head.Next
	}
	return false
}

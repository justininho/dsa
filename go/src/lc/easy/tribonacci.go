package easy

/*
(1137 | Easy) tribonacci
*/
func tribonacci(n int) int {
	if n == 0 {
		return 0
	} else if n <= 2 {
		return 1
	}
	t2, t1, t := 0, 1, 1
	for i := 3; i <= n; i++ {
		t2, t1, t = t1, t, t+t1+t2
	}
	return t
}

package easy

/*
(70 | Easy) Climbing Stairs
You are climbing a staircase. It takes n steps to reach the top.
Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
*/

func climbStairs(n int) int {
	memo := make(map[int]int)
	var climb func(int) int
	climb = func(n int) int {
		if n <= 1 {
			return 1
		} else if v, has := memo[n]; has {
			return v
		} else {
			memo[n] = climb(n-1) + climb(n-2)
			return memo[n]
		}
	}
	return climb(n)
}

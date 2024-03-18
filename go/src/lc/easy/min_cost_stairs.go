package easy

/*
(756 | Easy) minCostClimbingStairs
*/

func minCostClimbingStairs(cost []int) int {
	oneStep, twoStep := 0, 0
	for i := 2; i < len(cost)+1; i++ {
		twoStep, oneStep = oneStep, min(oneStep+cost[i-1], twoStep+cost[i-2])
	}
	return oneStep
}

//func minCostClimbingStairs(cost []int) int {
//	minCost := make([]int, len(cost)+1)
//	for i := 2; i < len(minCost); i++ {
//		oneStep := minCost[i-1] + cost[i-1]
//		twoStep := minCost[i-2] + cost[i-2]
//		minCost[i] = min(oneStep, twoStep)
//	}
//	return minCost[len(minCost)-1]
//}

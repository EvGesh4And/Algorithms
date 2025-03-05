package main

func minCostClimbingStairs(cost []int) int {

	a, b := cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		curr := cost[i] + min(a, b)
		a, b = b, curr
	}

	return min(a, b)
}

// or

func minCostClimbingStairs2(cost []int) int {
	a, b := 0, 0

	for i := 2; i < len(cost)+1; i++ {
		a, b = b, min(b+cost[i-1], a+cost[i-2])
	}
	return b
}

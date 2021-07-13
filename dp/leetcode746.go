package dp

func minCostClimbingStairs746(cost []int) int {
	if len(cost) < 2 {
		return 0
	} else {
		a, b := 0, 0
		for i := 2; i <= len(cost); i++ {
			if a+cost[i-2] <= b+cost[i-1] {
				a = a + cost[i-2]
				a, b = b, a
			} else {
				a = b + cost[i-1]
				a, b = b, a
			}
		}
		return b
	}
}

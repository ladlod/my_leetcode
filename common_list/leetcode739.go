package common_list

// 请根据每日 气温 列表 temperatures ，请计算在每一天需要等几天才会有更高的温度。如果气温在这之后都不会升高，请在该位置用 0 来代替。

// 思路：
// 栈，遍历入栈，每次出栈比遍历中数字小的所有数字，出栈数即为该位置的数

func DailyTemperatures(temperatures []int) []int {
	if len(temperatures) <= 0 {
		return temperatures
	}

	stack := make([]int, 0)
	stack = append(stack, 0)
	for i := 1; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			temperatures[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	for len(stack) > 0 {
		temperatures[stack[len(stack)-1]] = 0
		stack = stack[:len(stack)-1]
	}
	return temperatures
}

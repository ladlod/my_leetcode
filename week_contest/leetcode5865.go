package week_contest

import (
	"fmt"
	"math"
)

func firstDayBeenInAllRooms(nextVisit []int) int {
	visited := make(map[int]int)
	next := 0
	cnt := 0
	res := make([]int, len(nextVisit))
	for len(visited) < len(nextVisit) {
		cnt %= (int(math.Pow10(9)) + 7)
		cnt++
		if visited[next] == 0 {
			res[next] = cnt
		}
		visited[next]++
		if visited[next]&1 == 1 {
			cnt += res[next] - res[nextVisit[next]] + 1
			visited[next]++
		}
		next = (next + 1) % len(nextVisit)
	}
	fmt.Println(res)
	return res[len(res)-1] - 1
}

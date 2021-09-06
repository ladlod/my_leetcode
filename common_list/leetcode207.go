package common_list

// 你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。

// 在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。

// 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
// 请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/course-schedule
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：化简表达式 如 [1,2] [2,3] [3,4] 化简为{1,[4]}，若化简结果存在[x,x]，则不可能完成
//               如 [1,2] [2,3] [2,4] 则化简为{1,[3,4]}
// 以此思路，可以将表达式看成一张有向图，判断图是否存在环，存在环则不可能完成
func CanFinish(numCourses int, prerequisites [][]int) bool {
	preMap := make(map[int][]int) // 有向图
	for _, pre := range prerequisites {
		if _, ok := preMap[pre[0]]; !ok {
			preMap[pre[0]] = make([]int, 0)
		}
		preMap[pre[0]] = append(preMap[pre[0]], pre[1])
	}

	return !HasCicle(preMap)
}

func HasCicle(pic map[int][]int) bool {
	in := make(map[int]int) // 统计节点入度
	for _, vs := range pic {
		for _, v := range vs {
			in[v]++
		}
	}

	nin := make([]int, 0) // 记录入度为0的节点
	for k := range pic {
		if in[k] == 0 {
			nin = append(nin, k)
		}
	}

	for len(nin) > 0 {
		l := len(nin)
		for _, k := range nin {
			for _, v := range pic[k] {
				in[v]--
				if in[v] == 0 {
					nin = append(nin, v)
				}
			}
			delete(pic, k)
		}
		nin = nin[l:]
	}

	return len(pic) != 0
}

package common_list

type sortMultiNums [][]int

func (s sortMultiNums) Len() int {
	return len(s)
}

func (s sortMultiNums) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortMultiNums) Less(i, j int) bool {
	if len(s[i]) < 3 || len(s[j]) < 3 {
		return false
	}
	if s[i][0] == s[j][0] {
		return s[i][2] < s[j][2]
	}
	return s[i][0] < s[j][0]
}

func GetSkyline(buildings [][]int) [][]int {
	res := make([][]int, 0)

	return res
}

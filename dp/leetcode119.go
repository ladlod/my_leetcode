package dp

/*
	      [1]
	     [1, 1]
	    [1, 2, 1]
	   [1, 3, 3, 1]
	  [1, 4, 6, 4, 1]
	 [1, 5,10,10, 5, 1]
	[1, 6,15,20,15, 6, 1]
*/

func GetRowYanghui(rowIndex int) []int {
	var tmp, res []int
	for i := 0; i < rowIndex; i++ {
		res = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 {
				res[0] = 1
			} else if j == i {
				res[j] = tmp[j-1]
			} else {
				res[j] = tmp[j-1] + tmp[j]
			}
		}
		tmp = res
	}
	return res
}

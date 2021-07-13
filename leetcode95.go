package main

func generateTrees(n int) []*TreeNode {
	return generateTreesLR(1, n)
}

func generateTreesLR(l int, r int) []*TreeNode {
	res := make([]*TreeNode, 0)
	if l > r {
		return res
	} else if l == r {
		res = append(res, &TreeNode{
			Val: l,
		})
		return res
	} else {
		for i := l; i <= r; i++ {
			lefts := generateTreesLR(l, i-1)
			rights := generateTreesLR(i+1, r)
			if len(lefts) == 0 && len(rights) == 0 {
				continue
			} else if len(lefts) == 0 {
				for _, right := range rights {
					res = append(res, &TreeNode{
						Val:   i,
						Right: right,
					})
				}
			} else if len(rights) == 0 {
				for _, left := range lefts {
					res = append(res, &TreeNode{
						Val:  i,
						Left: left,
					})
				}
			} else {
				for _, left := range lefts {
					for _, right := range rights {
						res = append(res, &TreeNode{
							Val:   i,
							Left:  left,
							Right: right,
						})
					}
				}
			}
		}
	}
	return res
}

package week_contest

import "sort"

// 你正在参加一个多角色游戏，每个角色都有两个主要属性：攻击 和 防御 。给你一个二维整数数组 properties ，其中 properties[i] = [attacki, defensei] 表示游戏中第 i 个角色的属性。

// 如果存在一个其他角色的攻击和防御等级 都严格高于 该角色的攻击和防御等级，则认为该角色为 弱角色 。更正式地，如果认为角色 i 弱于 存在的另一个角色 j ，那么 attackj > attacki 且 defensej > defensei 。

// 返回 弱角色 的数量。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/the-number-of-weak-characters-in-the-game
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：
// 暴力：按照攻击值从低到高对给定角色排序，对每个角色，遍历攻击比它大的角色，判断是否存在防御也比他大的角色，如果存在则该角色为弱角色，时间复杂度O(n^2)
// 优化：按照攻击值从高到低对给定的角色排序，攻击值相同的角色按照防御从低到高排序，维护一个最高防御值，向后一次遍历，存在比最高防御值低的点即为弱角色
func numberOfWeakCharacters(properties [][]int) int {
	if len(properties) <= 0 {
		return 0
	}

	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] == properties[j][0] {
			return properties[i][1] < properties[j][1]
		}
		return properties[i][0] > properties[j][0]
	})

	res := 0
	maxD := properties[0][1]
	for i := 1; i < len(properties); i++ {
		if properties[i][1] < maxD {
			res++
		} else {
			maxD = properties[i][1]
		}
	}

	return res
}

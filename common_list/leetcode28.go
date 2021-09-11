package common_list

// 实现 strStr() 函数。

// 给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回  -1 。

// 说明：

// 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

// 对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与 C 语言的 strstr() 以及 Java 的 indexOf() 定义相符。

// 暴力做法
func StrStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

// kmp算法
// 前缀函数 记录一个字符串的[:i]部分最大前缀与后缀相等的长度为π(i)，如'aabaaab'的前缀函数为[0,1,0,1,2,2,3]
// 前缀函数计算
// π(i) = if str[π[i-1]+1] == str[i] : π(i-1) + 1
//		  else find a max j allow str[:j] == str[l-j:] => str[:j-1] == str[l-j-1:l-1] and str[j] == str[l] => π(j-1) and str[j] == str[l]
// eg: "issip#mississippi"
// 根据这个函数，我们可以将haystack append到needle后面，遍历到前缀函数值为len(needle)时，说明在haystack中出现了与needle相同的字符串
// 为了在新字符串中分割haystack和needle，将新字符串组成'needle#haystack'进行遍历
// 算法优化：无需真正的链接两个字符串，只需要按照needle # haystack的顺序遍历就行了
func StrStrKMP(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	pai := make([]int, 0)
	for i := 0; i < len(needle); i++ {
		if i == 0 {
			pai = append(pai, 0)
		} else {
			j := pai[i-1]
			for j > 0 && needle[i] != needle[j] {
				j = pai[j-1]
			}
			if needle[i] == needle[j] {
				j++
			}
			pai = append(pai, j)
		}
	}

	p := 0
	for i := 0; i < len(haystack); i++ {
		j := p
		for j > 0 && haystack[i] != needle[j] {
			j = pai[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		p = j
		if p == len(needle) {
			return i - p + 1
		}
	}

	return -1
}

// issi p #missi s sippi

// 0 0 0 1 0 0 0 1 2 3 4
// issip
// mississippi
// 0 1 2 3 4 2 3 4 5

package bst

// 不同的BST [leetcode 96]
func NumBSTs(n int) int {
	return bst_count(1, n)
}

func bst_count(low, hight int) int {
	if low >= hight {
		return 1
	}

	res := 0
	for i := low; i <= hight; i++ {
		left := bst_count(low, i-1)
		right := bst_count(i+1, hight)
		res += left * right
	}

	return res
}

// 卡塔兰数 C_nC [https://baike.baidu.com/item/catalan/7605685?fr=aladdin]
/**
* C0 = 1, Cn+1 = 2*(2n+1)/(n+2)*Cn
 */
// 使用数学方法
func catalan(n int) int {
	c := 1
	for i := 0; i < n; i++ {
		c = c * 2 * (2*i + 1) / (i + 2)
	}

	return c
}

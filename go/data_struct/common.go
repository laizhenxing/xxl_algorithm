package datastruct

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

// 获取两值的小值
func getMin(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

// 获取两值的大值
func getMax(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

// 绝对值
func abs(n1, n2 int) int {
	if n1 < n2 {
		n1 = n1 ^ n2
		n2 = n1 ^ n2
		n1 = n1 ^ n2
	}
	return n1 - n2
}

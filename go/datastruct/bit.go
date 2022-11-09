package datastruct

// 找寻一个数最右边第一个为1的位置
func FindRightIsOne(n int) int {
	notN := ^n + 1
	res := n & notN
	return res
}

// 交换 a,b 的值
// func Swap1(a, b int) {
// 	a = a ^ b
// 	b = a ^ b
// 	a = a ^ b
// }

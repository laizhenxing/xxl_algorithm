package bit

// 使用位运算做+ - * /

// +
func Addition(x, y int) int {
	sum := x
	for y != 0 {
		sum = x ^ y
		y = (x & y) << 1
		x = sum
	}
	return sum
}

// -
func Subtraction(x, y int) int {
	return Addition(x, oppositeNum(y))
}

func oppositeNum(n int) int {
	return Addition(^n, 1)
}

// *
func Multiplication(a, b int) int {
	return 0
}

// /
func Division(a, b int) int {
	return 0
}

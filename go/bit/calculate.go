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

// 取 n 的相反数
func oppositeNum(n int) int {
	return Addition(^n, 1)
}

// *
// 算法缺陷，当 b < 0 时，会在移位到-1时，形成死循环，因为-1的二进制是 11111111。
// 所以需要对 b < 0 进行讨论，令其必须大于0
func Multiplication(a, b int) int {
	if a < 0 && b < 0 {
		a, b = -a, -b
	} else if b < 0 {
		a, b = b, a
	}
	res := 0
	for b != 0 {
		if (b & 1) != 0 {
			res = Addition(res, a)
		}
		a <<= 1
		b >>= 1
	}
	return res
}

// /
func Division(a, b int) int {
	return 0
}

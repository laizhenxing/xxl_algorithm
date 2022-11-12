package bit

func Flip(n int) int {
	return n ^ 1
}

// 获取 n 的符号位
func Sign(n int) int {
	// 获取当前系统的位数
	bit := 32 << (^uint(0) >> 63)
	return Flip((n >> (bit-1)) & 1)
}

// 相反数
func OppositeNumber(n int) int {
	return ^n + 1
}

// 返回较大的值
func MaxByBit(a, b int) int {
	c := a - b
	sa := Sign(a)
	sb := Sign(b)
	sc := Sign(c)
	diffAb := sa ^ sb // 如果 a,b 符号不一样，则为1，否则为0
	sameAb := Flip(diffAb)
	reA := diffAb*sa + sameAb*sc
	reB := Flip(reA)
	return a*reA + b*reB
}

package datastruct

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

// 获取两值的小值
func Min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

// 获取两值的大值
func Max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

// 绝对值
func Abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}

func Swap[T int | int16 | int32 | int64 | float32 | float64](arr []T, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

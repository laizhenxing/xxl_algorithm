package datastruct

// 选择排序 O(N^2)
// 记住比i小的位置，最后再交换
func ChooseSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	length := len(nums)
	for i := 0; i < length-1; i++ {
		index := i
		for j := i + 1; j < length; j++ {
			if nums[j] < nums[index] {
				index = j
			}
		}
		if index != i {
			temp := nums[i]
			nums[i] = nums[index]
			nums[index] = temp
		}
	}
}

// 冒泡排序 O(N^2)
func BubbleSort(nums []int) {
	// 1. 0～N-1 大的往右换
	// 2. 0～N-2 大的往右换
	// ...
	if len(nums) <= 1 {
		return
	}

	for i := len(nums) - 1; i >= 0; i-- {
		for j := 0; j < i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j] = nums[j] ^ nums[j+1]
				nums[j+1] = nums[j] ^ nums[j+1]
				nums[j] = nums[j] ^ nums[j+1]
			}
		}
	}
}

// 插入排序：O(N^2)
func InsertSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	length := len(nums)
	i := 0
	for p := 1; p < length; p++ {
		temp := nums[p]
		for i = p; i > 0 && temp < nums[i-1]; i-- {
			nums[i] = nums[i-1]
		}
		nums[i] = temp
	}
}

// 快速排序 O(N*logN)
func QuickSort(nums []int) {

}

// 归并排序 O(N*logN)
func MergeSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	merge_process(nums, 0, len(nums)-1)
}

func merge_process(nums []int, l, r int) {
	if l == r {
		return
	}
	mid := l + ((r - l) >> 1)
	merge_process(nums, l, mid)
	merge_process(nums, mid+1, r)
	merge(nums, l, mid, r)
}

func merge(nums []int, l, mid, r int) {
	res := make([]int, r-l+1)
	index := 0
	p1, p2 := l, mid+1
	for ; (p1 <= l) && (p2 <= r); index++ {
		if nums[p1] < nums[p2] {
			res[index] = nums[p1]
			p1++
		} else {
			res[index] = nums[p2]
			p2++
		}
	}
	for p1 <= l {
		res[index] = nums[p1]
		p1++
		index++
	}
	for p2 <= r {
		res[index] = nums[p2]
		p2++
		index++
	}
	for i := 0; i < len(res); i++ {
		nums[l+i] = res[i]
	}
}

// 堆排序

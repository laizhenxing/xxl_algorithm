package dp

import (
	"algo/datastruct"
)

// 并查集
type Element[T comparable] struct {
	Value T
}

func NewElement[T comparable](value T) *Element[T] {
	ele := &Element[T]{Value: value}
	return ele
}

type UnionSet[T comparable] struct {
	elementMap map[T]Element[T]
	fatherMap  map[Element[T]]Element[T]
	sizeMap    map[Element[T]]int
}

func NewUnionSet[T comparable](list []T) *UnionSet[T] {
	elementMap := make(map[T]Element[T])
	fatherMap := make(map[Element[T]]Element[T])
	sizeMap := make(map[Element[T]]int)
	for _, v := range list {
		value := v
		ele := NewElement(v)
		elementMap[value] = *ele
		fatherMap[*ele] = *ele
		sizeMap[*ele] = 1
	}
	us := &UnionSet[T]{
		elementMap: elementMap,
		fatherMap:  fatherMap,
		sizeMap:    sizeMap,
	}
	return us
}

// 元素扁平化
func (us *UnionSet[T]) findHead(ele Element[T]) Element[T] {
	path := datastruct.NewStack()
	for e, ok := us.fatherMap[ele]; ok && ele != e; {
		path.Push(ele)
		ele = e
	}
	for !path.IsEmpty() {
		us.fatherMap[path.Pop().(Element[T])] = ele
	}
	return ele
}

// 普遍模版
func (us *UnionSet[T]) Find(ele Element[T]) Element[T] {
	if fa, ok := us.fatherMap[ele]; ok && fa != ele {
		us.fatherMap[ele] = us.Find(ele)
	}
	return us.fatherMap[ele]
}

func (us *UnionSet[T]) IsSameSet(a, b T) bool {
	ea, oka := us.elementMap[a]
	eb, okb := us.elementMap[b]
	if oka && okb {
		return us.findHead(ea) == us.findHead(eb)
	}
	return false
}

func (us *UnionSet[T]) Union(a, b T) {
	ea, oka := us.elementMap[a]
	eb, okb := us.elementMap[b]
	if oka && okb {
		af := us.findHead(ea)
		bf := us.findHead(eb)
		if af != bf {
			var big, small Element[T]
			if us.sizeMap[af] > us.sizeMap[bf] {
				big = af
				small = bf
			} else {
				big = bf
				small = af
			}
			us.fatherMap[small] = big
			// 改变 size
			us.sizeMap[big] = us.sizeMap[big] + us.sizeMap[small]
			// 从 sizeMap 移除被合并的节点信息，small
			delete(us.sizeMap, small)
		}
	}
}

// 最长连续序列
// func LongestConsecutive(nums []int) int {
// 	us := NewUnionSet[int](nums)
// 	ans := 0

// }

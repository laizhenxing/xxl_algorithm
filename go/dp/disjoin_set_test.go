package dp

import (
	"testing"
)

func TestUnion(t *testing.T) {

}

func TestIsSameSet(t *testing.T) {
	s := NewUnionSet[int]([]int{1, 2, 3, 4, 5, 6, 7, 8})
	str := NewUnionSet[string]([]string{"a", "b", "c", "d", "e"})

	type args struct {
		n1, n2 int
		s1, s2 string
	}
	intCases := []struct {
		name  string
		args  args
		wants bool
		union bool
	}{
		{
			name: "case1",
			args: args{
				n1: 1,
				n2: 2,
			},
			wants: false,
			union: false,
		},
		{
			name: "case2",
			args: args{
				n1: 1,
				n2: 2,
			},
			wants: true,
			union: true,
		},
		{
			name: "case3",
			args: args{
				n1: 2,
				n2: 3,
			},
			wants: false,
			union: false,
		},
		{
			name: "case4",
			args: args{
				n1: 2,
				n2: 3,
			},
			wants: true,
			union: true,
		},
		{
			name: "case5",
			args: args{
				n1: 1,
				n2: 3,
			},
			wants: true,
			union: false,
		},
		{
			name: "case6",
			args: args{
				n1: 1,
				n2: 8,
			},
			wants: false,
			union: false,
		},
	}

	strCases := []struct {
		name  string
		args  args
		wants bool
		union bool
	}{
		{
			name: "case7",
			args: args{
				s1: "a",
				s2: "b",
			},
			wants: false,
			union: false,
		},
		{
			name: "case8",
			args: args{
				s1: "a",
				s2: "b",
			},
			wants: true,
			union: true,
		},
		{
			name: "case9",
			args: args{
				s1: "a",
				s2: "c",
			},
			wants: false,
			union: false,
		},
		{
			name: "case10",
			args: args{
				s1: "a",
				s2: "c",
			},
			wants: true,
			union: true,
		},
		{
			name: "case11",
			args: args{
				s1: "a",
				s2: "d",
			},
			wants: false,
			union: false,
		},
		{
			name: "case12",
			args: args{
				s1: "a",
				s2: "e",
			},
			wants: false,
			union: false,
		},
	}

	for _, v := range intCases {
		t.Run(v.name, func(t *testing.T) {
			if v.union {
				s.Union(v.args.n1, v.args.n2)
			}

			if ok := s.IsSameSet(v.args.n1, v.args.n2); v.wants != ok {
				t.Fatalf("Except: %v; but got: %v", v.wants, ok)
			}
		})
	}

	for _, v := range strCases {
		t.Run(v.name, func(t *testing.T) {
			if v.union {
				str.Union(v.args.s1, v.args.s2)
			}

			if ok := str.IsSameSet(v.args.s1, v.args.s2); v.wants != ok {
				t.Fatalf("Except: %v; but got: %v", v.wants, ok)
			}
		})
	}
}

// 最长连续序列
// Leetcode 128 https://leetcode.cn/problems/longest-consecutive-sequence/description/
func TestLongestConsecutive(t *testing.T) {
	
}
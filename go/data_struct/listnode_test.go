package datastruct

import (
	"reflect"
	"testing"
)

func TestSingleNodeAdd(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		args args
		want struct {
			length int
			vals   []int
		}
	}{
		// TODO: Add test cases.
		{
			name: "Add/one",
			args: args{
				val: 1,
			},
			want: struct {
				length int
				vals   []int
			}{
				length: 1,
				vals:   []int{1},
			},
		},
		{
			name: "Add/two",
			args: args{
				val: 2,
			},
			want: struct {
				length int
				vals   []int
			}{
				length: 2,
				vals:   []int{1, 2},
			},
		},
	}

	singleNode := &SingleNode{}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			singleNode.Add(test.args.val)
			singleNode.Display()
			if length, vals := singleNode.GetDetails(); length != test.want.length || !reflect.DeepEqual(vals, test.want.vals) {
				t.Fatalf("Add val to SingleNode, Except length: %v, vals: %v; but got length: %v, vals: %v",
					test.want.length,
					test.want.vals,
					length,
					vals)
			}
		})
	}
}

func TestSingleNodeMulAdd(t *testing.T) {
	type args struct {
		vals []int
	}
	cases := []struct {
		name  string
		args  args
		wants []int
	}{
		{
			name: "case1",
			args: args{
				vals: []int{1, 0, 9, 4, 3, 6, 8, 7, 5},
			},
			wants: []int{0, 1, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "case2",
			args: args{
				vals: []int{11, 10, 93, 41, 32, 64, 86, 72, 58},
			},
			wants: []int{10, 11, 32, 41, 58, 64, 72, 86, 93},
		},
	}

	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			singleNode := &SingleNode{}
			singleNode.MulAdd(ca.args.vals)
			if _, vals := singleNode.GetDetails(); !reflect.DeepEqual(vals, ca.wants) {
				t.Fatalf("Add multi vals to Node Fail. Except: %v; but got %v", ca.wants, vals)
			}
		})
	}
}

func TestSingleNodeReverse(t *testing.T) {
	type args struct {
		vals []int
	}

	cases := []struct {
		name  string
		args  args
		wants []int
	}{
		{
			name: "case1",
			args: args{
				[]int{1, 0, 9, 4, 3, 6, 8, 7, 5},
			},
			wants: []int{9, 8, 7, 6, 5, 4, 3, 1, 0},
		},
		{
			name: "case2",
			args: args{
				vals: []int{11, 10, 93, 41, 32, 64, 86, 72, 58},
			},
			wants: []int{93, 86, 72, 64, 58, 41, 32, 11, 10},
		},
	}

	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			singleNode := &SingleNode{}
			singleNode.MulAdd(ca.args.vals)
			singleNode = ReverseListWithHead(singleNode)
			if _, vals := singleNode.GetDetails(); !reflect.DeepEqual(ca.wants, vals) {
				t.Fatalf("Reverse list fail. Except vals: %v, but got vals: %v", ca.wants, vals)
			}
		})
	}
}

func TestSingleNodeFindCommPart(t *testing.T) {
	type args struct {
		node1Vals []int
		node2Vals []int
	}
	cases := []struct {
		name  string
		args  args
		wants []int
	}{
		{
			name: "case1",
			args: args{
				node1Vals: []int{1, 3, 5, 7, 8, 9, 10},
				node2Vals: []int{2, 3, 4, 5, 6, 8, 10},
			},
			wants: []int{3, 5, 8, 10},
		},
		{
			name: "case2",
			args: args{
				node1Vals: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				node2Vals: []int{2, 3, 4, 5, 6, 8, 10},
			},
			wants: []int{2, 3, 4, 5, 6, 8, 10},
		},
	}

	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			n1, n2 := &SingleNode{}, &SingleNode{}
			n1.MulAdd(ca.args.node1Vals)
			n2.MulAdd(ca.args.node2Vals)
			if comm := SingleNodeFindCommPart(n1, n2); !reflect.DeepEqual(comm, ca.wants) {
				t.Fatalf("FindCommPart Fail.Except: %v, but got: %v", ca.wants, comm)
			}
		})
	}
}

func TestIsPalindrome1(t *testing.T) {
	type args struct {
		vals []int
	}
	cases := []struct {
		name  string
		args  args
		wants bool
	}{
		{
			name: "case1",
			args: args{
				vals: []int{1, 2, 2, 1},
			},
			wants: true,
		},
		{
			name: "case2",
			args: args{
				vals: []int{1, 2, 3, 2, 1},
			},
			wants: true,
		},
		{
			name: "case3",
			args: args{
				vals: []int{1, 2, 3, 2, 1, 0},
			},
			wants: false,
		},
		{
			name: "case4",
			args: args{
				vals: []int{},
			},
			wants: true,
		},
	}

	for _, ca := range cases {
		singleNode := CreateLinkWithVals(ca.args.vals)
		t.Run(ca.name, func(t *testing.T) {
			if isPal := IsPalindrome1(singleNode); isPal != ca.wants {
				t.Fatalf("%v is palindrome? Except: %v, but got: %v", ca.args.vals, ca.wants, isPal)
			}
		})
	}
}

func TestIsPalindrome2(t *testing.T) {
	type args struct {
		vals []int
	}
	cases := []struct {
		name  string
		args  args
		wants bool
	}{
		{
			name: "case1",
			args: args{
				vals: []int{1, 2, 2, 1},
			},
			wants: true,
		},
		{
			name: "case2",
			args: args{
				vals: []int{1, 2, 3, 2, 1},
			},
			wants: true,
		},
		{
			name: "case3",
			args: args{
				vals: []int{1, 2, 3, 2, 1, 0},
			},
			wants: false,
		},
		{
			name: "case4",
			args: args{
				vals: []int{},
			},
			wants: true,
		},
	}

	for _, ca := range cases {
		singleNode := CreateLinkWithVals(ca.args.vals)
		t.Run(ca.name, func(t *testing.T) {
			if isPal := IsPalindrome2(singleNode); isPal != ca.wants {
				t.Fatalf("%v is palindrome? Except: %v, but got: %v", ca.args.vals, ca.wants, isPal)
			}
		})
	}
}

func TestIsPalindrome3(t *testing.T) {
	type args struct {
		vals []int
	}
	cases := []struct {
		name  string
		args  args
		wants bool
	}{
		{
			name: "case1",
			args: args{
				vals: []int{1, 2, 2, 1},
			},
			wants: true,
		},
		{
			name: "case2",
			args: args{
				vals: []int{1, 2, 3, 2, 1},
			},
			wants: true,
		},
		{
			name: "case3",
			args: args{
				vals: []int{1, 2, 3, 2, 1, 0},
			},
			wants: false,
		},
		{
			name: "case4",
			args: args{
				vals: []int{},
			},
			wants: true,
		},
	}

	for _, ca := range cases {
		singleNode := CreateLinkWithVals(ca.args.vals)
		t.Run(ca.name, func(t *testing.T) {
			if isPal := IsPalindrome3(singleNode); isPal != ca.wants {
				t.Fatalf("%v is palindrome? Except: %v, but got: %v", ca.args.vals, ca.wants, isPal)
			}
		})
	}
}

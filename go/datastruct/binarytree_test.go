package datastruct

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeserializeByString(t *testing.T) {
	root := DeserializeByString("1_2_4_#_8_#_#_5_9_#_#_#_3_6_#_#_7_#_#_")
	PreOrderByRecursive(root)
	str := SerializeByPre(root)
	fmt.Println(str)
}

func TestSerializeByPre(t *testing.T) {
	type args struct {
		root *BinaryTree
	}

	cases := []struct {
		name  string
		args  args
		wants string
	}{
		{
			name: "case1",
			args: args{
				root: DeserializeByString("1_2_4_#_8_#_#_5_9_#_#_#_3_6_#_#_7_#_#_"),
			},
			wants: "1_2_4_#_8_#_#_5_9_#_#_#_3_6_#_#_7_#_#_",
		},
		{
			name: "case2",
			args: args{
				root: DeserializeByString("1_2_#_4_#_5_#_6_#_#_3_#_#_"),
			},
			wants: "1_2_#_4_#_5_#_6_#_#_3_#_#_",
		},
	}

	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			if nodeStr := SerializeByPre(ca.args.root); nodeStr != ca.wants {
				t.Fatalf("serialize binary tree. Except: %v, but got: %v", ca.wants, nodeStr)
			}
		})
	}
}

// 递归先序遍历
func TestPreOrderByRecursive(t *testing.T) {
	type args struct {
		root *BinaryTree
	}
	cases := []struct {
		name  string
		args  args
		wants []int
	}{
		{
			name: "case1",
			args: args{
				root: DeserializeByString("1_2_4_#_8_#_#_5_9_#_#_#_3_6_#_#_7_#_#_"),
			},
			wants: []int{1, 2, 4, 8, 5, 9, 3, 6, 7},
		},
		{
			name: "case2",
			args: args{
				root: DeserializeByString("1_2_#_4_#_5_#_6_#_#_3_#_#_"),
			},
			wants: []int{1, 2, 4, 5, 6, 3},
		},
		{
			name: "case3",
			args: args{
				root: DeserializeByString("7_6_4_8_#_#_#_3_#_#_5_2_#_#_1_#_#_"),
			},
			wants: []int{7, 6, 4, 8, 3, 5, 2, 1},
		},
	}

	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			if vals := PreOrderByRecursive(ca.args.root); !reflect.DeepEqual(ca.wants, vals) {
				t.Fatalf("pre_order by recursive fail. Except: %v, but got: %v", ca.wants, vals)
			}
		})
	}
}

// 非递归先序遍历
func TestPreOrderByUnRecursive(t *testing.T) {
	type args struct {
		root *BinaryTree
	}
	cases := []struct {
		name  string
		args  args
		wants []int
	}{
		{
			name: "case1",
			args: args{
				root: DeserializeByString("1_2_4_#_8_#_#_5_9_#_#_#_3_6_#_#_7_#_#_"),
			},
			wants: []int{1, 2, 4, 8, 5, 9, 3, 6, 7},
		},
		{
			name: "case2",
			args: args{
				root: DeserializeByString("1_2_#_4_#_5_#_6_#_#_3_#_#_"),
			},
			wants: []int{1, 2, 4, 5, 6, 3},
		},
		{
			name: "case3",
			args: args{
				root: DeserializeByString("7_6_4_8_#_#_#_3_#_#_5_2_#_#_1_#_#_"),
			},
			wants: []int{7, 6, 4, 8, 3, 5, 2, 1},
		},
	}

	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			if vals := PreOrderByUnRecursive(ca.args.root); !reflect.DeepEqual(ca.wants, vals) {
				t.Fatalf("pre_order by unrecursive fail. Except: %v, but got: %v", ca.wants, vals)
			}
		})
	}
}

// 中序遍历递归
func TestInOrderByRecursive(t *testing.T) {
	type args struct {
		root *BinaryTree
	}
	cases := []struct {
		name  string
		args  args
		wants []int
	}{
		{
			name: "case1",
			args: args{
				root: DeserializeByString("1_2_4_#_8_#_#_5_9_#_#_#_3_6_#_#_7_#_#_"),
			},
			wants: []int{4, 8, 2, 9, 5, 1, 6, 3, 7},
		},
		{
			name: "case2",
			args: args{
				root: DeserializeByString("1_2_#_4_#_5_#_6_#_#_3_#_#_"),
			},
			wants: []int{2, 4, 5, 6, 1, 3},
		},
		{
			name: "case3",
			args: args{
				root: DeserializeByString("7_6_4_8_#_#_#_3_#_#_5_2_#_#_1_#_#_"),
			},
			wants: []int{8, 4, 6, 3, 7, 2, 5, 1},
		},
	}

	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			if vals := InOrderByRecursive(ca.args.root); !reflect.DeepEqual(ca.wants, vals) {
				t.Fatalf("id_order by unrecursive fail. Except: %v, but got: %v", ca.wants, vals)
			}
		})
	}
}

// 非递归中序遍历
func TestInOrderByUnRecursive(t *testing.T) {
	type args struct {
		root *BinaryTree
	}
	cases := []struct {
		name  string
		args  args
		wants []int
	}{
		{
			name: "case1",
			args: args{
				root: DeserializeByString("1_2_4_#_8_#_#_5_9_#_#_#_3_6_#_#_7_#_#_"),
			},
			wants: []int{4, 8, 2, 9, 5, 1, 6, 3, 7},
		},
		{
			name: "case2",
			args: args{
				root: DeserializeByString("1_2_#_4_#_5_#_6_#_#_3_#_#_"),
			},
			wants: []int{2, 4, 5, 6, 1, 3},
		},
		{
			name: "case3",
			args: args{
				root: DeserializeByString("7_6_4_8_#_#_#_3_#_#_5_2_#_#_1_#_#_"),
			},
			wants: []int{8, 4, 6, 3, 7, 2, 5, 1},
		},
	}

	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			if vals := InOrderByUnRecursive(ca.args.root); !reflect.DeepEqual(ca.wants, vals) {
				t.Fatalf("id_order by unrecursive fail. Except: %v, but got: %v", ca.wants, vals)
			}
		})
	}
}

// 递归后续遍历
func TestPostOrderByRecursive(t *testing.T) {
	type args struct {
		root *BinaryTree
	}
	cases := []struct {
		name  string
		args  args
		wants []int
	}{
		{
			name: "case1",
			args: args{
				root: DeserializeByString("1_2_4_#_8_#_#_5_9_#_#_#_3_6_#_#_7_#_#_"),
			},
			wants: []int{8, 4, 9, 5, 2, 6, 7, 3, 1},
		},
		{
			name: "case2",
			args: args{
				root: DeserializeByString("1_2_#_4_#_5_#_6_#_#_3_#_#_"),
			},
			wants: []int{6, 5, 4, 2, 3, 1},
		},
		{
			name: "case3",
			args: args{
				root: DeserializeByString("7_6_4_8_#_#_#_3_#_#_5_2_#_#_1_#_#_"),
			},
			wants: []int{8, 4, 3, 6, 2, 1, 5, 7},
		},
	}

	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			if vals := PostOrderByRecursive(ca.args.root); !reflect.DeepEqual(ca.wants, vals) {
				t.Fatalf("post_order by recursive fail. Except: %v, but got: %v", ca.wants, vals)
			}
		})
	}
}

// 非递归后续遍历
func TestPostOrderByUnRecursive(t *testing.T) {
	type args struct {
		root *BinaryTree
	}
	cases := []struct {
		name  string
		args  args
		wants []int
	}{
		{
			name: "case1",
			args: args{
				root: DeserializeByString("1_2_4_#_8_#_#_5_9_#_#_#_3_6_#_#_7_#_#_"),
			},
			wants: []int{8, 4, 9, 5, 2, 6, 7, 3, 1},
		},
		{
			name: "case2",
			args: args{
				root: DeserializeByString("1_2_#_4_#_5_#_6_#_#_3_#_#_"),
			},
			wants: []int{6, 5, 4, 2, 3, 1},
		},
		{
			name: "case3",
			args: args{
				root: DeserializeByString("7_6_4_8_#_#_#_3_#_#_5_2_#_#_1_#_#_"),
			},
			wants: []int{8, 4, 3, 6, 2, 1, 5, 7},
		},
	}

	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			if vals := PostOrderByUnRecursive(ca.args.root); !reflect.DeepEqual(ca.wants, vals) {
				t.Fatalf("post_order by unrecursive fail. Except: %v, but got: %v", ca.wants, vals)
			}
		})
	}
}

// 判断二叉树是否是搜索树（BST）
func TestIsBST(t *testing.T) {
	type args struct {
		root *BinaryTree
	}
	cases := []struct {
		name  string
		args  args
		wants bool
	}{
		{
			name:  "case1",
			args:  args{root: &BinaryTree{}},
			wants: true,
		},
		{
			name:  "case2",
			args:  args{root: DeserializeByString("6_4_#_#_8_7_#_#_9_#_#_")},
			wants: true,
		},
		{
			name:  "case3",
			args:  args{root: DeserializeByString("6_4_2_1_#_#_3_#_#_5_#_#_8_7_#_#_9_#_#_")},
			wants: true,
		},
		{
			name:  "case4",
			args:  args{root: DeserializeByString("8_6_5_#_#_7_#_#_9_7_#_#_11_#_#_")},
			wants: false,
		},
		{
			name:  "case5",
			args:  args{root: DeserializeByString("2_1_#_#_3_#_#_")},
			wants: true,
		},
		{
			name:  "case6",
			args:  args{root: DeserializeByString("2_#_#_")},
			wants: true,
		},
	}

	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			if isBST := IsBST(ca.args.root); ca.wants != isBST {
				t.Fatalf("binary tree is BST? Except: %v; but got: %v", ca.wants, isBST)
			}
		})
	}
}

// 判断二叉树是否是搜索树（BST）
func TestIsBST2(t *testing.T) {
	type args struct {
		root *BinaryTree
	}
	cases := []struct {
		name  string
		args  args
		wants bool
	}{
		{
			name:  "case1",
			args:  args{root: &BinaryTree{}},
			wants: true,
		},
		{
			name:  "case2",
			args:  args{root: DeserializeByString("6_4_#_#_8_7_#_#_9_#_#_")},
			wants: true,
		},
		{
			name:  "case3",
			args:  args{root: DeserializeByString("6_4_2_1_#_#_3_#_#_5_#_#_8_7_#_#_9_#_#_")},
			wants: true,
		},
		{
			name:  "case4",
			args:  args{root: DeserializeByString("8_6_5_#_#_7_#_#_9_7_#_#_11_#_#_")},
			wants: false,
		},
		{
			name:  "case5",
			args:  args{root: DeserializeByString("2_1_#_#_3_#_#_")},
			wants: true,
		},
		{
			name:  "case6",
			args:  args{root: DeserializeByString("2_#_#_")},
			wants: true,
		},
	}

	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			if isBST := IsBST2(ca.args.root); ca.wants != isBST {
				t.Fatalf("binary tree is BST? Except: %v; but got: %v", ca.wants, isBST)
			}
		})
	}
}

// 判断二叉树是否是搜索树（BST）
func TestIsBST3(t *testing.T) {
	type args struct {
		root *BinaryTree
	}
	cases := []struct {
		name  string
		args  args
		wants bool
	}{
		{
			name:  "case1",
			args:  args{root: &BinaryTree{}},
			wants: true,
		},
		{
			name:  "case2",
			args:  args{root: DeserializeByString("6_4_#_#_8_7_#_#_9_#_#_")},
			wants: true,
		},
		{
			name:  "case3",
			args:  args{root: DeserializeByString("6_4_2_1_#_#_3_#_#_5_#_#_8_7_#_#_9_#_#_")},
			wants: true,
		},
		{
			name:  "case4",
			args:  args{root: DeserializeByString("8_6_5_#_#_7_#_#_9_7_#_#_11_#_#_")},
			wants: false,
		},
		{
			name:  "case5",
			args:  args{root: DeserializeByString("2_1_#_#_3_#_#_")},
			wants: true,
		},
		{
			name:  "case6",
			args:  args{root: DeserializeByString("2_#_#_")},
			wants: true,
		},
	}

	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			if isBST := IsBST3(ca.args.root); ca.wants != isBST {
				t.Fatalf("binary tree is BST? Except: %v; but got: %v", ca.wants, isBST)
			}
		})
	}
}

func TestIsFullTree(t *testing.T) {
	type args struct {
		root *BinaryTree
	}
	cases := []struct {
		name  string
		args  args
		wants bool
	}{
		{
			name:  "case1",
			args:  args{root: &BinaryTree{}},
			wants: true,
		},
		{
			name:  "case2",
			args:  args{root: DeserializeByString("6_4_#_#_8_7_#_#_9_#_#_")},
			wants: false,
		},
		{
			name:  "case3",
			args:  args{root: DeserializeByString("6_4_2_1_#_#_3_#_#_5_#_#_8_7_#_#_9_#_#_")},
			wants: false,
		},
		{
			name:  "case4",
			args:  args{root: DeserializeByString("8_6_5_#_#_7_#_#_9_7_#_#_11_#_#_")},
			wants: true,
		},
		{
			name:  "case5",
			args:  args{root: DeserializeByString("2_1_#_#_3_#_#_")},
			wants: true,
		},
		{
			name:  "case6",
			args:  args{root: DeserializeByString("2_#_#_")},
			wants: true,
		},
	}

	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			if isBST := IsFullTree(ca.args.root); ca.wants != isBST {
				t.Fatalf("binary tree is full tree? Except: %v; but got: %v", ca.wants, isBST)
			}
		})
	}
}

func TestIsBalance(t *testing.T) {
	type args struct {
		root *BinaryTree
	}
	cases := []struct {
		name  string
		args  args
		wants bool
	}{
		{
			name:  "case1",
			args:  args{root: &BinaryTree{}},
			wants: true,
		},
		{
			name:  "case2",
			args:  args{root: DeserializeByString("6_4_#_#_8_7_#_#_9_#_#_")},
			wants: true,
		},
		{
			name:  "case3",
			args:  args{root: DeserializeByString("6_4_2_1_#_#_3_#_#_5_#_#_8_7_#_#_9_#_#_")},
			wants: true,
		},
		{
			name:  "case4",
			args:  args{root: DeserializeByString("8_6_5_#_#_7_#_#_9_7_#_#_11_#_#_")},
			wants: true,
		},
		{
			name:  "case5",
			args:  args{root: DeserializeByString("2_1_#_#_3_#_#_")},
			wants: true,
		},
		{
			name:  "case6",
			args:  args{root: DeserializeByString("2_#_#_")},
			wants: true,
		},
		{
			name:  "case7",
			args:  args{root: DeserializeByString("6_4_2_1_0_#_#_10_#_#_3_#_#_5_#_#_8_7_#_#_9_#_#_")},
			wants: false,
		},
	}

	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			if isBST := IsBalanceTree(ca.args.root); ca.wants != isBST {
				t.Fatalf("binary tree is balance? Except: %v; but got: %v", ca.wants, isBST)
			}
		})
	}
}

// 宽度优先遍历
func TestBreadthFirstTraversal(t *testing.T) {
	type args struct {
		root *BinaryTree
	}
	cases := []struct {
		name  string
		args  args
		wants []int
	}{
		{
			name: "case1",
			args: args{
				root: DeserializeByString("1_2_4_#_8_#_#_5_9_#_#_#_3_6_#_#_7_#_#_"),
			},
			wants: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "case2",
			args: args{
				root: DeserializeByString("1_2_#_4_#_5_#_6_#_#_3_#_#_"),
			},
			wants: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "case3",
			args: args{
				root: DeserializeByString("7_6_4_8_#_#_#_3_#_#_5_2_#_#_1_#_#_"),
			},
			wants: []int{7, 6, 5, 4, 3, 2, 1, 8},
		},
	}

	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			if vals := BreadthFirstTraversal(ca.args.root); !reflect.DeepEqual(ca.wants, vals) {
				t.Fatalf("post_order by unrecursive fail. Except: %v, but got: %v", ca.wants, vals)
			}
		})
	}
}

// 宽度优先遍历
func TestIsCBT(t *testing.T) {
	type args struct {
		root *BinaryTree
	}
	cases := []struct {
		name  string
		args  args
		wants bool
	}{
		{
			name: "case1",
			args: args{
				root: DeserializeByString("1_2_4_#_8_#_#_5_9_#_#_#_3_6_#_#_7_#_#_"),
			},
			wants: false,
		},
		{
			name: "case2",
			args: args{
				root: DeserializeByString("1_2_#_4_#_5_#_6_#_#_3_#_#_"),
			},
			wants: false,
		},
		{
			name: "case3",
			args: args{
				root: DeserializeByString("7_6_4_8_#_#_#_3_#_#_5_2_#_#_1_#_#_"),
			},
			wants: true,
		},
		{
			name: "case4",
			args: args{
				root: DeserializeByString("1_2_4_8_#_#_9_#_#_5_10_#_#_#_3_6_#_#_7_#_#_"),
			},
			wants: true,
		},
	}

	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			if vals := IsCBT(ca.args.root); !reflect.DeepEqual(ca.wants, vals) {
				t.Fatalf("post_order by unrecursive fail. Except: %v, but got: %v", ca.wants, vals)
			}
		})
	}
}

func TestLowestAncestor2(t *testing.T) {
	type args struct {
		root  *BinaryTree
		node1 *BinaryTree
		node2 *BinaryTree
	}
	cases := []struct {
		name  string
		args  args
		wants []int
	}{
		{
			name: "case1",
			args: args{
				root:  DeserializeByString("1_2_4_#_8_#_#_5_9_#_#_#_3_6_#_#_7_#_#_"),
				node1: nil,
				node2: nil,
			},
			wants: nil,
		},
		{
			name: "case2",
			args: args{
				root:  DeserializeByString("1_2_#_4_#_5_#_6_#_#_3_#_#_"),
				node1: nil,
				node2: nil,
			},
			wants: nil,
		},
		{
			name: "case3",
			args: args{
				root:  DeserializeByString("7_6_4_8_#_#_#_3_#_#_5_2_#_#_1_#_#_"),
				node1: nil,
				node2: nil,
			},
			wants: nil,
		},
		{
			name: "case4",
			args: args{
				root:  DeserializeByString("1_2_4_8_#_#_9_#_#_5_10_#_#_#_3_6_#_#_7_#_#_"),
				node1: nil,
				node2: nil,
			},
			wants: nil,
		},
	}

	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			if vals := IsCBT(ca.args.root); !reflect.DeepEqual(ca.wants, vals) {
				t.Fatalf("post_order by unrecursive fail. Except: %v, but got: %v", ca.wants, vals)
			}
		})
	}
}

package datastruct

import (
	"fmt"
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

func TestPreOrderByNonRecursive(t *testing.T) {
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
			wants: "",
		},
		{
			name: "case2",
			args: args{
				root: DeserializeByString("1_2_#_4_#_5_#_6_#_#_3_#_#_"),
			},
			wants: "1_2_#_4_#_5_#_6_#_#_3_#_#_",
		},
		{
			name: "case3",
			args: args{
				root: DeserializeByString("7_6_4_8_#_#_#_3_#_#_5_2_#_#_1_#_#_"),
			},
			wants: "7_6_4_8_#_#_#_3_#_#_5_2_#_#_1_#_#_",
		},
	}

	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			PreOrderByUnRecursive(ca.args.root)
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
		wants string
	}{
		{
			name: "case1",
			args: args{
				root: DeserializeByString("1_2_4_#_8_#_#_5_9_#_#_#_3_6_#_#_7_#_#_"),
			},
			wants: "",
		},
		{
			name: "case2",
			args: args{
				root: DeserializeByString("1_2_#_4_#_5_#_6_#_#_3_#_#_"),
			},
			wants: "1_2_#_4_#_5_#_6_#_#_3_#_#_",
		},
		{
			name: "case3",
			args: args{
				root: DeserializeByString("7_6_4_8_#_#_#_3_#_#_5_2_#_#_1_#_#_"),
			},
			wants: "7_6_4_8_#_#_#_3_#_#_5_2_#_#_1_#_#_",
		},
	}

	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			InOrderByRecursive(ca.args.root)
		})
	}
}

func TestInOrderByUnRecursive(t *testing.T) {
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
			wants: "",
		},
		{
			name: "case2",
			args: args{
				root: DeserializeByString("1_2_#_4_#_5_#_6_#_#_3_#_#_"),
			},
			wants: "1_2_#_4_#_5_#_6_#_#_3_#_#_",
		},
		{
			name: "case3",
			args: args{
				root: DeserializeByString("7_6_4_8_#_#_#_3_#_#_5_2_#_#_1_#_#_"),
			},
			wants: "7_6_4_8_#_#_#_3_#_#_5_2_#_#_1_#_#_",
		},
	}

	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			InOrderByUnRecursive(ca.args.root)
		})
	}
}

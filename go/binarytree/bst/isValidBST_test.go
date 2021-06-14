package bst

import (
	"algo/binarytree"
	"testing"
)

func TestIsValidBST(t *testing.T) {
	type args struct {
		root *binarytree.TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "is BST",
			args: args{
				root: binarytree.Deserialize("6,5,#,#,15,10,#,#,20,#,#,"),
			},
			want: true,
		},
		{
			name: "is not BST",
			args: args{
				root: binarytree.Deserialize("10,5,#,#,15,6,#,#,20,#,#,"),
			},
			want: false,
		},
		{
			name: "is not BST",
			args: args{
				root: binarytree.Deserialize("2,2,#,#,2,#,#,"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidBST(tt.args.root); got != tt.want {
				t.Errorf("IsValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *binarytree.TreeNode
		min  *binarytree.TreeNode
		max  *binarytree.TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root, tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

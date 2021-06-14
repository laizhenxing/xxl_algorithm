package bst

import (
	"algo/binarytree"
	"testing"
)

func TestIsInBST(t *testing.T) {
	type args struct {
		root   *binarytree.TreeNode
		target int
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
			if got := IsInBST(tt.args.root, tt.args.target); got != tt.want {
				t.Errorf("IsInBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

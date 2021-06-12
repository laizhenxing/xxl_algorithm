package bst

import (
	"algo/binarytree"
	"testing"
)

func TestKthSmallest(t *testing.T) {
	type args struct {
		root *binarytree.TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KthSmallest(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("KthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_kth_traverse(t *testing.T) {
	type args struct {
		root *binarytree.TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log()
			kth_traverse(tt.args.root, tt.args.k)
		})
	}
}

package bst

import (
	"algo/binarytree"
	"reflect"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	type args struct {
		root *binarytree.TreeNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *binarytree.TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteNode(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMinNode(t *testing.T) {
	type args struct {
		node *binarytree.TreeNode
	}
	tests := []struct {
		name string
		args args
		want *binarytree.TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinNode(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMinNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMaxNode(t *testing.T) {
	type args struct {
		node *binarytree.TreeNode
	}
	tests := []struct {
		name string
		args args
		want *binarytree.TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxNode(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMaxNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

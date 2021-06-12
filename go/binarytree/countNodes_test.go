package binarytree

import (
	"testing"
)

func Test_countNodes(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "nodes=4",
			args: args{
				root: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val:   1,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val:   4,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
				},
			},
			want: 4,
		},
		{
			name: "node=3",
			args: args{
				root: &TreeNode{
					Val:   2,
					Left:  &TreeNode{},
					Right: &TreeNode{},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountNodes(tt.args.root); got != tt.want {
				t.Errorf("CountNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountLeaves(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "leaves=0",
			args: args{
				root: nil,
			},
			want: 0,
		},
		{
			name: "leaves=2",
			args: args{
				root: &TreeNode{
					Val:   2,
					Left:  &TreeNode{},
					Right: &TreeNode{},
				},
			},
			want: 2,
		},
		{
			name: "leaves=3",
			args: args{
				root: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{},
						Right: &TreeNode{},
					},
					Right: &TreeNode{},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountLeaves(tt.args.root); got != tt.want {
				t.Errorf("CountLeaves() = %v, want %v", got, tt.want)
			}
		})
	}
}

package binarytree

import (
	"reflect"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{
			name: "Both In",
			args: args{
				root: Deserialize("3,5,6,#,#,2,7,#,#,4,#,#,1,0,#,#,8,#,#,"),
				p:    Deserialize("5,6,#,#,2,7,#,#,4,#,#,"),
				q:    Deserialize("1,0,#,#,8,#,#,"),
				// p: Deserialize("3,5,6,#,#,2,7,#,#,4,#,#,1,0,#,#,8,#,#,"),
				// q: Deserialize("3,5,6,#,#,2,7,#,#,4,#,#,1,0,#,#,8,#,#,"),
			},
			want: Deserialize("3,5,6,#,#,2,7,#,#,4,#,#,1,0,#,#,8,#,#,"),
		},
		{
			name: "Both In",
			args: args{
				root: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
				},
				p: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				q: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
			want: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
		},
		{
			name: "null",
			args: args{
				root: nil,
				p: &TreeNode{
					Val:   1,
					Left:  &TreeNode{},
					Right: &TreeNode{},
				},
				q: &TreeNode{
					Val:   2,
					Left:  &TreeNode{},
					Right: &TreeNode{},
				},
			},
			want: nil,
		},
		{
			name: "Both Not",
			args: args{
				root: &TreeNode{
					Val:   3,
					Left:  &TreeNode{},
					Right: &TreeNode{},
				},
				p: &TreeNode{
					Val:   1,
					Left:  &TreeNode{},
					Right: &TreeNode{},
				},
				q: &TreeNode{
					Val:   2,
					Left:  &TreeNode{},
					Right: &TreeNode{},
				},
			},
			want: nil,
		},
		{
			name: "P In",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   1,
						Left:  &TreeNode{},
						Right: &TreeNode{},
					},
					Right: &TreeNode{},
				},
				p: &TreeNode{
					Val:   1,
					Left:  &TreeNode{},
					Right: &TreeNode{},
				},
				q: &TreeNode{
					Val:   2,
					Left:  &TreeNode{},
					Right: &TreeNode{},
				},
			},
			want: &TreeNode{
				Val:   1,
				Left:  &TreeNode{},
				Right: &TreeNode{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}

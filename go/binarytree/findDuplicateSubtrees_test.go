package binarytree

import (
	"reflect"
	"testing"
)

func TestFindDuplicateSubtrees(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		// TODO: Add test cases.
		{
			name: "nums=2",
			args: args{
				root: new(Codec).Deserialize("1,2,4,#,#,#,3,2,4,#,#,#,4,#,#,"),
			},
			want: []*TreeNode{
				&TreeNode{Val: 4},
				&TreeNode{Val: 2, Left: &TreeNode{Val: 4}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindDuplicateSubtrees(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindDuplicateSubtrees() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_traverse(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := traverse(tt.args.root); got != tt.want {
				t.Errorf("traverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

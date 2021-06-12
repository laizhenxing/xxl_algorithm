package bst

import (
	"algo/binarytree"
	"reflect"
	"testing"
)

func TestConvertBST(t *testing.T) {
	type args struct {
		root *binarytree.TreeNode
	}
	tests := []struct {
		name string
		args args
		want *binarytree.TreeNode
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				root: binarytree.Deserialize("4,1,0,#,#,2,#,3,#,#,6,5,#,#,7,#,8,#,#,"),
			},
			want: binarytree.Deserialize("30,36,36,#,#,35,#,33,#,#,21,26,#,#,15,#,8,#,#,"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertBST(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_conver_traverse(t *testing.T) {
	type args struct {
		root *binarytree.TreeNode
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
			conver_traverse(tt.args.root)
		})
	}
}

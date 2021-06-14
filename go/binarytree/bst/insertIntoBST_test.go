package bst

import (
	"algo/binarytree"
	"reflect"
	"testing"
)

func TestInsertIntoBST(t *testing.T) {
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
		{
			name: "success",
			args: args{
				root: binarytree.Deserialize("4,2,1,#,#,3,#,#,7,#,#,"),
				val:  5,
			},
			want: binarytree.Deserialize("4,2,1,#,#,3,#,#,7,5,#,#,#,"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertIntoBST(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertIntoBST() = %v, want %v", binarytree.Serialize(got), binarytree.Serialize(tt.want))
			}
		})
	}
}

package binarytree

import "testing"

func TestPrintf(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	c := new(Codec)
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "t",
			args: args{
				c.Deserialize("7,4,3,1,0,#,#,#,#,2,#,#,3,#,#,"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Printf(tt.args.root)
		})
	}
}

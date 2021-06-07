package binarytree

import (
	"testing"
)

func TestInvertTreePreorder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name   string
		args   args
		expect string
	}{
		// TODO: Add test cases.
		{
			name: "true",
			args: args{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:  2,
						Left: nil,
						Right: &TreeNode{
							Val:   4,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
			},
			expect: "1,3,#,#,2,4,#,#,#,",
		},
		{
			name: "false",
			args: args{
				&TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 5,
						},
					},
					Right: &TreeNode{
						Val: 6,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 7,
							Right: &TreeNode{
								Val: 3,
							},
						},
					},
				},
			},
			expect: "4,6,7,3,#,#,#,3,#,#,8,#,5,#,#,",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := InvertTreePreorder(tt.args.root)
			codec := new(Codec)
			actual := codec.Serialize(r)

			if tt.name == "true" {
				if tt.expect != actual {
					t.Log("expect: ", tt.expect)
					t.Log("actual: ", actual)
					t.Fatal("invert binary tree fail")
				}
			}

			if tt.name == "false" {
				if tt.expect != actual {
					t.Log("expect: ", tt.expect)
					t.Log("actual: ", actual)
					t.Fatal("invert binary tree fail")
				}
			}
		})
	}
}

func TestInvertTreePostorder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name   string
		args   args
		expect string
	}{
		// TODO: Add test cases.
		{
			name: "true",
			args: args{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:  2,
						Left: nil,
						Right: &TreeNode{
							Val:   4,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
			},
			expect: "1,3,#,#,2,4,#,#,#,",
		},
		{
			name: "false",
			args: args{
				&TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 5,
						},
					},
					Right: &TreeNode{
						Val: 6,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 7,
							Right: &TreeNode{
								Val: 3,
							},
						},
					},
				},
			},
			expect: "4,6,7,3,#,#,#,3,#,#,8,#,5,#,#,",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := InvertTreePostorder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("InvertTreePostorder() = %v, want %v", got, tt.want)
			// }
			r := InvertTreePreorder(tt.args.root)
			codec := new(Codec)
			actual := codec.Serialize(r)

			if tt.name == "true" {
				if tt.expect != actual {
					t.Log("expect: ", tt.expect)
					t.Log("actual: ", actual)
					t.Fatal("invert binary tree fail")
				}
			}

			if tt.name == "false" {
				if tt.expect != actual {
					t.Log("expect: ", tt.expect)
					t.Log("actual: ", actual)
					t.Fatal("invert binary tree fail")
				}
			}
		})
	}
}

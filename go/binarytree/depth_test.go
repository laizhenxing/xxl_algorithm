package binarytree

import (
	"fmt"
	"testing"
)

func TestDepthPreoder(t *testing.T) {
	type args struct {
		root  *TreeNode
		depth int
	}
	c := new(Codec)
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "4",
			args: args{
				root:  c.Deserialize("4,6,7,3,#,#,#,3,#,#,8,#,5,#,#,"),
				depth: 0,
			},
			want: 4,
		},
		{
			name: "5",
			args: args{
				root:  c.Deserialize("7,4,3,1,0,#,#,#,#,2,#,#,3,#,#,"),
				depth: 0,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		fmt.Printf("%#v\n", c.Deserialize("7,4,3,1,0,#,#,#,#,2,#,#,3,#,#,"))
		t.Run(tt.name, func(t *testing.T) {
			if got := DepthPreoder(tt.args.root, tt.args.depth); got != tt.want {
				t.Errorf("DepthPreoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDepthInorder(t *testing.T) {
	type args struct {
		root  *TreeNode
		depth int
	}
	c := new(Codec)
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "4",
			args: args{
				root:  c.Deserialize("4,6,7,3,#,#,#,3,#,#,8,#,5,#,#,"),
				depth: 0,
			},
			want: 4,
		},
		{
			name: "5",
			args: args{
				root:  c.Deserialize("7,4,3,1,0,#,#,#,#,2,#,#,3,#,#,"),
				depth: 0,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DepthInorder(tt.args.root, tt.args.depth); got != tt.want {
				t.Errorf("DepthInorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDepthPostoder(t *testing.T) {
	type args struct {
		root  *TreeNode
		depth int
	}
	c := new(Codec)
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "4",
			args: args{
				root:  c.Deserialize("4,6,7,3,#,#,#,3,#,#,8,#,5,#,#,"),
				depth: 0,
			},
			want: 4,
		},
		{
			name: "5",
			args: args{
				root:  c.Deserialize("7,4,3,1,0,#,#,#,#,2,#,#,3,#,#,"),
				depth: 0,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DepthPostoder(tt.args.root, tt.args.depth); got != tt.want {
				t.Errorf("DepthPostoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

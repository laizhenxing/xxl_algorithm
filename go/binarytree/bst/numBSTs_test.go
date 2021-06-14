package bst

import "testing"

func TestNumBSTs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "equal",
			args: args{
				n: 3,
			},
			want: 5,
		},
		{
			name: "equal",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "equal",
			args: args{
				n: 0,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumBSTs(tt.args.n); got != tt.want {
				t.Errorf("NumBSTs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bst_count(t *testing.T) {
	type args struct {
		low   int
		hight int
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
			if got := bst_count(tt.args.low, tt.args.hight); got != tt.want {
				t.Errorf("bst_count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_catalan(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "equal",
			args: args{
				n: 3,
			},
			want: 5,
		},
		{
			name: "equal",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "equal",
			args: args{
				n: 0,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := catalan(tt.args.n); got != tt.want {
				t.Errorf("catalan() = %v, want %v", got, tt.want)
			}
		})
	}
}

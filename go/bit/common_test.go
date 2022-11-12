package bit

import "testing"

func TestMaxByBit(t *testing.T) {
	type args struct {
		a, b int
	}
	cases := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{89, 23},
			want: 89,
		},
		{
			name: "case2",
			args: args{23, 893},
			want: 893,
		},
		{
			name: "case3",
			args: args{-89, 23},
			want: 23,
		},
		{
			name: "case4",
			args: args{89, -23},
			want: 89,
		},
		{
			name: "case5",
			args: args{-89, -23},
			want: -23,
		},
	}

	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			if res := MaxByBit(ca.args.a, ca.args.b); ca.want != res {
				t.Fatalf("Excepcted: %v, but got: %v", ca.want, res)
			}
		})
	}
}

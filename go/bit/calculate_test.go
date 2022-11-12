package bit

import "testing"

func TestAddition(t *testing.T) {
	type args struct {
		x, y int
	}
	cases := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{4, 5},
			want: 9,
		},
		{
			name: "case2",
			args: args{5, -3},
			want: 2,
		},
		{
			name: "case3",
			args: args{-5, 2},
			want: -3,
		},
	}

	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			if res := Addition(ca.args.x, ca.args.y); ca.want != res {
				t.Fatalf("Excepted: %v, but got: %v", ca.want, res)
			}
		})
	}
}

func TestSubtraction(t *testing.T) {
	type args struct {
		x, y int
	}
	cases := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{4, 5},
			want: -1,
		},
		{
			name: "case2",
			args: args{5, -3},
			want: 8,
		},
		{
			name: "case3",
			args: args{-5, 2},
			want: -7,
		},
	}

	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			if res := Subtraction(ca.args.x, ca.args.y); ca.want != res {
				t.Fatalf("Excepted: %v, but got: %v", ca.want, res)
			}
		})
	}
}

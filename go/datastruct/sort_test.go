package datastruct

import "testing"

func TestBubbleSort(t *testing.T) {

}

func TestArs(t *testing.T) {
	Ars()
}

func TestIntervalSchedule(t *testing.T) {
	type args struct {
		ints [][]int
	}
	cases := []struct {
		name  string
		args  args
		wants int
	}{
		{
			name: "case1",
			args: args{
				ints: [][]int{
					{1, 2},
					{2, 4},
					{3, 6},
				},
			},
			wants: 2,
		},
		{
			name:  "case2",
			args:  args{ints: [][]int{{1, 3}, {3, 5}, {4, 7}, {8, 10}, {5, 9}}},
			wants: 3,
		},
	}

	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			if count := IntervalSchedule(ca.args.ints); ca.wants != count {
				t.Fatalf("Except: %v, but got: %v", ca.wants, count)
			}
		})
	}
}

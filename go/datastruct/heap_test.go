package datastruct

import (
	"reflect"
	"testing"
)

func TestHeapify(t *testing.T) {
	type args struct {
		arr             []int
		index, heapSize int
	}

	cases := []struct {
		name  string
		args  args
		wants []int
	}{
		{
			name: "case1",
			args: args{
				arr:      []int{7, 6, 5, 4, 3, 2, 1},
				index:    0,
				heapSize: 7,
			},
			wants: []int{7, 6, 5, 4, 3, 2, 1},
		},
		{
			name: "case2",
			args: args{
				arr:      []int{1},
				index:    0,
				heapSize: 1,
			},
			wants: []int{1},
		},
		{
			name: "case3",
			args: args{
				arr:      []int{1, 2, 3},
				index:    0,
				heapSize: 3,
			},
			wants: []int{3, 2, 1},
		},
	}

	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			Heapify(ca.args.arr, ca.args.index, ca.args.heapSize)
			if !reflect.DeepEqual(ca.wants, ca.args.arr) {
				t.Fatalf("Except: %v; but got: %v", ca.wants, ca.args.arr)
			}
		})
	}
}

func TestHeapSort(t *testing.T) {
	type args struct {
		arr             []int
		index, heapSize int
	}

	cases := []struct {
		name  string
		args  args
		wants []int
	}{
		{
			name: "case1",
			args: args{
				arr:      []int{3, 5, 8, 4, 6, 7, 0},
				index:    0,
				heapSize: 7,
			},
			wants: []int{0, 3, 4, 5, 6, 7, 8},
		},
		{
			name: "case2",
			args: args{
				arr:      []int{1},
				index:    0,
				heapSize: 1,
			},
			wants: []int{1},
		},
		{
			name: "case3",
			args: args{
				arr:      []int{3, 1, 2},
				index:    0,
				heapSize: 3,
			},
			wants: []int{1, 2, 3},
		},
	}

	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			HeapSort(ca.args.arr)
			if !reflect.DeepEqual(ca.wants, ca.args.arr) {
				t.Fatalf("Except: %v; but got: %v", ca.wants, ca.args.arr)
			}
		})
	}
}

func TestPQ(t *testing.T) {
	PQ()
}

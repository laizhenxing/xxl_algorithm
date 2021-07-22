package lru

import (
	"reflect"
	"sync"
	"testing"
)

func TestLRUCache_Get(t *testing.T) {
	type args struct {
		key interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
		{
			name: "Get/one",
			args: args{
				key: "one",
			},
			want: 1,
		},
		{
			name: "Get/two",
			args: args{
				key: "two",
			},
			want: 2,
		},
	}
	lru := NewLRUCache(3)
	lru.Put("one", 1)
	lru.Put("two", 2)
	lru.Put("three", 3)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lru.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LRUCache.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLRUCache_Put(t *testing.T) {
	type fields struct {
		keyMap *sync.Map
		cache  *DoubleList
		cap    int
	}
	type args struct {
		key interface{}
		val interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lru := &LRUCache{
				keyMap: tt.fields.keyMap,
				cache:  tt.fields.cache,
				cap:    tt.fields.cap,
			}
			lru.Put(tt.args.key, tt.args.val)
		})
	}
}

func TestLRUCache_First(t *testing.T) {
	tests := []struct {
		name string
		want map[interface{}]interface{}
	}{
		// TODO: Add test cases.
		{
			name: "get/one",
			want: map[interface{}]interface{}{
				"one": 1,
			},
		},
		{
			name: "get/two",
			want: map[interface{}]interface{}{
				"two": 2,
			},
		},
		{
			name: "get/three",
			want: map[interface{}]interface{}{
				"three": 3,
			},
		},
	}

	keys := []string{"one", "two", "three"}

	lru := NewLRUCache(3)
	lru.Put("one", 1)
	lru.Put("two", 2)
	lru.Put("three", 3)
	for i, tt := range tests {
		lru.Get(keys[i])
		t.Run(tt.name, func(t *testing.T) {
			if got := lru.First(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LRUCache.First() = %v, want %v", got, tt.want)
			}
		})
	}
}

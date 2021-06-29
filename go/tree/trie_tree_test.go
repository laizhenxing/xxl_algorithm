package tree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewTrie(t *testing.T) {
	type args struct {
		fullPath string
		path     string
		children map[rune]*Trie
	}
	tests := []struct {
		name string
		args args
		want *Trie
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTrie(tt.args.fullPath, tt.args.path, tt.args.children); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTrie() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrie_Search(t *testing.T) {
	trie := NewTrie("/", "/", make(map[rune]*Trie))
	// trie.Insert("/a/b/c")
	trie.Insert("/a")
	fmt.Println("trie", trie)

	type fields struct {
		fullPath string
		path     string
		children map[rune]*Trie
		isEnd    bool
	}
	type args struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Trie
	}{
		// TODO: Add test cases.
		{
			name: "success",
			fields: fields{
				fullPath: "",
				path:     "",
				children: map[rune]*Trie{
					rune('/'): {
						fullPath: "/a",
						path:     "/",
						children: map[rune]*Trie{
							rune('a'): {
								fullPath: "/a",
								path:     "/a",
								children: map[rune]*Trie{},
								isEnd:    true,
							},
						},
						isEnd: true,
					},
				},
				isEnd: false,
			},
			args: args{
				path: "/a",
			},
			want: &Trie{
				fullPath: "/a",
				path:     "/a",
				children: make(map[rune]*Trie),
				isEnd:    true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Trie{
				fullPath: tt.fields.fullPath,
				path:     tt.fields.path,
				children: tt.fields.children,
				isEnd:    tt.fields.isEnd,
			}

			if got := tr.Search(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Trie.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrie_Insert(t *testing.T) {
	type fields struct {
		fullPath string
		path     string
		children map[rune]*Trie
		isEnd    bool
	}
	type args struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "success",
			fields: fields{
				fullPath: "",
				path:     "",
				children: map[rune]*Trie{
					0: {
						fullPath: "",
						path:     "",
						children: map[rune]*Trie{
							0: {},
						},
						isEnd: false,
					},
				},
				isEnd: false,
			},
			args: args{
				path: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Trie{
				fullPath: tt.fields.fullPath,
				path:     tt.fields.path,
				children: tt.fields.children,
				isEnd:    tt.fields.isEnd,
			}
			tr.Insert(tt.args.path)
		})
	}
}

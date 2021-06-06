package binarytree

import (
	"testing"
)

func TestSerialize(t *testing.T) {
	bt := &TreeNode{
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
	}

	expect := "1,2,#,4,#,#,3,#,#,"

	codec := &Codec{}
	actual := codec.Serialize(bt)

	if expect != actual {
		t.Log("actual: ", actual)
		t.Log("expect: ", expect)
		t.Fatal("Fail to serialize")
	}
}

func TestDeserialize(t *testing.T) {
	expect := "1,2,#,4,#,#,3,#,#,"

	codec := &Codec{}

	node := codec.Deserialize(expect)

	actual := codec.Serialize(node)

	if expect != actual {
		t.Log("actual: ", actual)
		t.Log("expect: ", expect)
		t.Fatal("Fail to deserialize")
	}
}

#### BST 查找某个节点的遍历框架
```go
func BST(root *TreeNode, int target) {
    if (root ==  nil) {
        return
    }
    if (root.Val == target) {
        // TODO
    }

    if (root.Val < target) {
        BST(root.Right, target)
    }
    if (root.Val > target) {
        BST(root.Left, target)
    }
}
```

#### BST 删除一个一个数
```go
func DeleteNode(root *TreeNode, val int) *TreeNode {
    if root == nil { return nil }
    if root.Val == val {
        // 找到，删除
    } else if ( root.Val > val ) {
        root.Left = DeleteNode(root.Left, val)
    } else if ( root.Val < val ) {
        root.Right = DeleteNode(root.Right, val)
    }

    return root
}
```
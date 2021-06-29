package tree

// 最小前缀树
type Trie struct {
	fullPath string
	path     string
	children map[rune]*Trie
	isEnd    bool
}

func NewTrie(fullPath, path string, children map[rune]*Trie) *Trie {
	return &Trie{
		fullPath: fullPath,
		path:     path,
		children: children,
	}
}

func (t *Trie) Search(path string) *Trie {
	node := t
	for _, ch := range path {
		if node.children[ch] == nil {
			return nil
		}

		node = node.children[ch]
	}

	return node
}

func (t *Trie) Insert(path string) {
	node := t
	for i, ch := range path {
		if node.children[ch] == nil {
			children := make(map[rune]*Trie)
			node.children[ch] = NewTrie(path, path[:i+1], children)
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie) Count() int {
	node := t
	count := 0
	if node == nil || node.isEnd {
		return count
	}
	for _, child := range node.children {
		if child.isEnd {
			count += 1
		}
		count += child.Count()
	}

	return count
}

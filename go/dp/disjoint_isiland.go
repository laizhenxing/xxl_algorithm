package dp

// 计算岛的数量
// leetcode 200 https://leetcode.cn/problems/number-of-islands/description/
func NumIsIlands(grid [][]byte) int {
	if grid == nil || grid[0] == nil {
		return 0
	}
	iu := NewIlandUnion(grid)
	n := len(grid)
	m := len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				if i-1 >= 0 && grid[i-1][j] == '1' {
					iu.Union(i*m+j, (i-1)*m+j)
				}
				if i+1 < n && grid[i+1][j] == '1' {
					iu.Union(i*m+j, (i+1)*m+j)
				}
				if j-1 >= 0 && grid[i][j-1] == '1' {
					iu.Union(i*m+j, i*m+j-1)
				}
				if j+1 < m && grid[i][j+1] == '1' {
					iu.Union(i*m+j, i*m+j+1)
				}
			}
		}
	}

	return iu.Count()
}

type IlandUnion struct {
	parent []int
	rank   []int
	count  int
}

// 初始化集合
func NewIlandUnion(grid [][]byte) *IlandUnion {
	n := len(grid)
	m := len(grid[0])
	count := 0
	parent := make([]int, n*m)
	rank := make([]int, n*m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				parent[i*m+j] = i*m + j
				count++
			}
		}
	}

	return &IlandUnion{
		parent: parent,
		rank:   rank,
		count:  count,
	}
}

func (iu *IlandUnion) findHead(x int) int {
	if x != iu.parent[x] {
		iu.parent[x] = iu.findHead(iu.parent[x])
	}
	return iu.parent[x]
}

func (iu *IlandUnion) Union(x, y int) {
	xRoot := iu.findHead(x)
	yRoot := iu.findHead(y)
	if xRoot != yRoot {
		if iu.rank[xRoot] >= iu.rank[yRoot] {
			iu.parent[yRoot] = xRoot
			iu.rank[xRoot] += 1
		} else {
			iu.parent[xRoot] = yRoot
			iu.rank[yRoot] += 1
		}
		iu.count--
	}
}

func (iu *IlandUnion) Count() int {
	return iu.count
}

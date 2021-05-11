type Union struct {
	parent []int;
	rank []int;
	count int; // 集合の数
  }
  
  func newUnion(grid [][]byte) *Union {
	u := new(Union)
	m := len(grid)
	n := len(grid[0])
	u.parent = make([]int, 0)
	u.rank = make([]int, 0)
	
	for i := 0; i < m; i++ {
	  for j := 0; j < n; j++ {
		if grid[i][j] == '1' { 
		  u.parent = append(u.parent, i*n+j) // 1だったら、自分の位置をrootにする
		  u.count++
		} else {
		  u.parent = append(u.parent, -1) // 0だったら、木の要素ではないので-1を設定
		}
		u.rank = append(u.rank, 0) // rankは最初全て0
	  }
	}
	return u
  }
  
  // 木のrootを探す
  func (u *Union) find(i int) int {
	// 自分自身が親だったらrootと判定できる
	if u.parent[i] != i { // rootじゃなかったら
	  u.parent[i] = u.find(u.parent[i])  // 経路圧縮
	}
	return u.parent[i]
  }
  
  // 結合する
  func (u *Union) union(x int, y int) {
	
	// それぞれのrootを探す
	rootx := u.find(x)
	rooty := u.find(y)
	
	// rootが同じなら、既に同じ木に属している
	if rootx == rooty {
	  return
	}
	
	if u.rank[rootx] < u.rank[rooty] {
	  u.parent[rootx] = rooty
	} else {
	  u.parent[rooty] = rootx
	  if u.rank[rootx] == u.rank[rooty] {
		u.rank[rootx]++
	  }
	}
	
	// 1の集合がグルーピングされたので、集合の数を減らす
	u.count--
  }
  
  func (u *Union) getCount() int {
	return u.count
  }
  
  func numIslands(grid [][]byte) int {
	nr := len(grid)
	if nr == 0 {
	  return 0
	}
	nc := len(grid[0])
	
	uf := newUnion(grid)
	
	for r := 0; r < nr; r++ {
	  for c := 0; c < nc; c++ {
		
		if grid[r][c] == '1' {
		  grid[r][c] = '0'
			   if r-1 >= 0 && grid[r-1][c] == '1' { // 一つ上のノードを結合する
			uf.union(r*nc+c, (r-1)*nc+c)
		  }
		  if r+1 < nr && grid[r+1][c] == '1' {
			uf.union(r*nc+c, (r+1)*nc+c)
		  }
		  if c-1 >= 0 && grid[r][c-1] == '1' {
			uf.union(r*nc+c, r*nc+c-1)
		  }
		  if c+1 < nc && grid[r][c+1] == '1' {
			uf.union(r*nc+c, r*nc+c+1)
		  }
		}
		
	  }
	}
	
	return uf.getCount()
  }
  
  
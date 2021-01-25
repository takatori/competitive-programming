type Neighbor struct {
	r int;
	c int;
  }
  
  func numIslands(grid [][]byte) int {
	nr := len(grid)
	
	if nr == 0 {
	  return 0
	}
	
	nc := len(grid[0])
	
	num_islands := 0
	
	for r := 0; r < nr; r++ {
	  for c := 0; c < nc; c++ {
		if grid[r][c] == '1' {
		  num_islands++
		  grid[r][c] = '0' // mark as visited
		  q := make([]Neighbor, 0)
		  q = append(q, Neighbor{r,c}) // push
		  
		  for len(q) > 0 {
			rc := q[0]
			q = q[1:] // pop
			row := rc.r
			col := rc.c
  
			if row-1 >= 0 && grid[row-1][col] == '1' {
			   q = append(q, Neighbor{row-1, col})
			   grid[row-1][col] = '0'
			}
  
			if row+1 < nr && grid[row+1][col] == '1' {
			  q = append(q, Neighbor{row+1, col})
			  grid[row+1][col] = '0'
			}
  
			if col-1 >= 0 && grid[row][col-1] == '1' {
			  q = append(q, Neighbor{row, col-1})
			  grid[row][col-1] = '0'
			}
  
			if col+1 < nc && grid[row][col+1] == '1' {
			  q = append(q, Neighbor{row, col+1})
			  grid[row][col+1] = '0'
			}
		  }
		}
	  }
	}
	
	return num_islands
  }
  
  
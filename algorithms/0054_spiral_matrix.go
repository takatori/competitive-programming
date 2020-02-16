func spiralOrder(matrix [][]int) []int {
    
    if len(matrix) == 0 {
        return []int{}
    }
    
    ans := make([]int, 0)
    
    r1, r2 := 0, len(matrix)-1
    c1, c2 := 0, len(matrix[0])-1
    
    for r1 <= r2 && c1 <= c2 {
        for _, v := range spiralCoords(r1,c1,r2,c2) {
            ans = append(ans, matrix[v[0]][v[1]])
        }
        r1 += 1
        r2 -= 1
        c1 += 1
        c2 -= 1
    }

    return ans
}

func spiralCoords(r1, c1, r2, c2 int) [][]int {
    
    ans := make([][]int, 0)
    
    for c := c1; c < c2+1; c++ {
        ans = append(ans, []int{r1, c})
    }
    
    for r := r1+1; r < r2+1; r++ {
        ans = append(ans, []int{r, c2})
    }
    
    if r1 < r2 && c1 < c2 {
        for c := c2-1; c > c1; c-- {
            ans = append(ans, []int{r2, c})
        }
        for r := r2; r > r1; r-- {
            ans = append(ans, []int{r, c1})
        }
    }
    
    return ans
}

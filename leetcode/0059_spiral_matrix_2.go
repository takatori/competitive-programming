var count int

func generateMatrix(n int) [][]int {
    
    result := make([][]int, n)
    for i := 0; i < n; i++ {
        result[i] = make([]int, n)
    }
    
    r1, r2 := 0, n-1
    c1, c2 := 0, n-1
    count = 1
    
    for r1 <= r2 && c1 <= c2 {
        fillSpiralCoords(r1,c1,r2,c2, result)
        r1 += 1
        r2 -= 1
        c1 += 1
        c2 -= 1
    }

    return result    
}

func fillSpiralCoords(r1, c1, r2, c2 int, result [][]int) {
    
    for c := c1; c < c2+1; c++ {
        result[r1][c] = count
        count++
    }
    
    for r := r1+1; r < r2+1; r++ {
        result[r][c2] = count
        count++
    }

    for c := c2-1; c > c1; c-- {
        result[r2][c] = count
        count++
    }

    for r := r2; r > r1; r-- {
        result[r][c1] = count
        count++
    }

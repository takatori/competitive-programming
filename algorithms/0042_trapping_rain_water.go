func trap(height []int) int { // [4,2,3]
   
    var result int
    var i int
    
    for i < len(height) { // i = 1,  len(height) = 3
        j := i+1 // j = 2
        w := 0
        for j < len(height) && height[i] > height[j] { // 2 > 3
            w += height[i] - height[j] // w = 2 + 1
            j++ // j = 3
        }
        if j < len(height) {
            result += w // 1 + 4 + 1
            i = j // i = 10
        } else {
            i++
        }
    }
    
    return result
}

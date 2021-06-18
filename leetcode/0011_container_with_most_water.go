func maxArea(height []int) int {
    var max int
    
    for i := 0; i < len(height) - 1; i++ {
        for j := i+1; j < len(height); j++ {
            var h int
            if height[i] > height[j] {
                h = height[j]
            } else {
                h = height[i]
            }
            area := h * (j - i)
            if area > max {
                max = area
            }            
        }
    }    
    return max
}

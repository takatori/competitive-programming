func trap(height []int) int { // [4,2,3]
   
    if len(height) == 0 {
        return 0
    }
    
    var result int
    size := len(height)
    left_max := make([]int, size)
    right_max := make([]int, size)
    
    left_max[0] = height[0]
    for i := 1; i < size; i++ {
        left_max[i] = int(math.Max(float64(height[i]), float64(left_max[i-1])))
    }
    
    right_max[size-1] = height[size-1]
    for i := size-2; i >= 0; i-- {
        right_max[i] = int(math.Max(float64(height[i]), float64(right_max[i+1])))
    }
    
    for i := 1; i < size-1; i++ {        
        result += int(math.Min(float64(left_max[i]), float64(right_max[i]))) - height[i]
    }
    
    return result
}

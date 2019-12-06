
func reverse(x int) int {
    
    target := x
    result := 0
    keep := false
    
    if x < 0 {
        target = -x
    }

    for i := target; i > 0; i = i / 10 {
        n := i % 10
        if n != 0 || keep {
            keep = true
            result = result * 10 + n
        } 
    }
    
    if result < -2147483648 || result > 2147483647 {
        return 0
    } else if x < 0 {
        return -result
    } else {
        return result
    }
    
}

func divide(dividend int, divisor int) int {
    
    var result int
    
    flag := true
    
    if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
        flag = false
    }
    
    if dividend < 0 {
        dividend = -dividend
    }
    
    if divisor < 0 {
        divisor = -divisor
    }
    
    
    for dividend >= divisor {
        dividend = dividend - divisor
        result++
    }
    
    if flag && result > math.MaxInt32 {
        return math.MaxInt32
    }
    
    if flag {
        return result
    } else {
        return -result
    }
}

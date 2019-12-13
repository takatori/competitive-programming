func isPalindrome(x int) bool {
    
    var rev int
    
    if x < 0 {
        return false
    }
    
    tmp := x
    
    for tmp > 0 {        
        rev = rev * 10 + tmp % 10
        tmp = tmp / 10
    }
    
    return x == rev
}

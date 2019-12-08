func longestPalindrome(s string) string {
    
    length := len(s)
    maxI := 0
    maxJ := 0
    max := 0
    
    for i := 0; i < length; i++ {
        for j := i + 1;  j <= length; j++ {
            if isPalindrome(s[i:j]) && max < j - i {
                max = j - i
                maxI = i
                maxJ = j
            }
        }
    }
    
    return s[maxI:maxJ]
}

func isPalindrome(s string) bool {
    
    length := len(s)
    
    if length < 2 {
        return true
    }
    
    if s[0] == s[length-1] {
        return isPalindrome(s[1:length-1])
    } 
    
    return false
    
}

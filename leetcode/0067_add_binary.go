func addBinary(a string, b string) string {
    
    result := make([]int, 0)
    carry := 0
    i,j := len(a)-1, len(b)-1
    
    for i >= 0 || j >= 0 {
        var ai, bi int
        if i >= 0 {
            ai = toInt(a[i])
        }
        if j >= 0 {
            bi = toInt(b[j])
        }
        result = append(result, (ai+bi+carry)%2)
        carry = (ai+bi+carry)/2        
        i--
        j--
    }
    
    ans := ""
    
    for _, n := range result {
        ans = strconv.Itoa(n) + ans
    }
    if carry == 1 {
        ans = "1" + ans
    }
    return ans
    
}


func toInt(a byte) int {
    return int(a - '0')
}

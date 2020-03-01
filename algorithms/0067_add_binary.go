func addBinary(a string, b string) string {
    
    result := 0
    i := len(a)-1
    j := len(b)-1
    carry := 0
    
    for i >= 0 || j >= 0 {
        var ai, bi int
        if i >= 0 {
            ai = toInt(a[i])
        }
        if j >= 0 {
            bi = toInt(b[j])
        }
        result = ((ai+bi+carry)%2)*10 + result
        carry = (ai+bi+carry)/2        
        i--
        j--
    }
    return strconv.Itoa(result)
    
}


func toInt(a byte) int {
    return int(a - '0')
}

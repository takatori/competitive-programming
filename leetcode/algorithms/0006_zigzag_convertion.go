func convert(s string, numRows int) string {
    
    dict := make([][]byte, numRows)
    m := 2 * numRows - 2
    if numRows == 1 {
        return s
    }
        
    for i, c := range s {
        n := i % m
        if n >= numRows {
            n = m - n
        }
        dict[n] = append(dict[n], byte(c))
    }
    
    var result []byte
    
    for _, r := range dict {
        result = append(result, r...)
    }
    
    return string(result)
    
}

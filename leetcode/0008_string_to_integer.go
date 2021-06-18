func myAtoi(str string) int {
    
    var result int
    sign := 1
    prev := ' '
    
    for _, c := range str {
        if c == ' ' || c == '+' || c == '-'{
            if prev != ' ' {
                return sign * result
            }
            if c == '-' {
                sign = -1
            }
            prev = c
        } else if isDigit(c) {
            tmp := 10 * result + int(c - '0')
            if sign < 0 && tmp > 2147483648 {
                return -2147483648
            } else if sign > 0 && tmp > 2147483647 {
                return 2147483647
            }
            prev = c
            result = tmp
        } else {
            return sign * result
        }
    }
    return sign * result
}

func isDigit(c rune) bool {
    return '0' <= c && c <= '9'
}

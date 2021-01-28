import "strconv"

var memo map[int]int

func numDecodings(s string) int {
  memo = make(map[int]int)
  return recursiveWithMemo(0, s)
}

func recursiveWithMemo(index int, str string) int {
  if v, ok := memo[index]; ok {
    return v
  }
  
  if index == len(str) {
    return 1
  }
  
  if str[index] == '0' {
    return 0
  }
  
  if index == len(str)-1 {
    return 1
  }
  
  ans := recursiveWithMemo(index+1, str)
  i, _ := strconv.Atoi(str[index:index+2]) 
  if i <= 26 {
    ans += recursiveWithMemo(index+2, str)
  }
  
  memo[index] = ans
  
  return ans
}


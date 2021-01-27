
var count int
var memo map[string]int

func numDecodings(s string) int {
  
  if len(s) == 0 {
    return 0
  }
  memo = make(map[string]int, 0)
  f(s)
  
  return count
}

func f(s string) {
  if len(s) == 0 {
    count++
    return
  }
  
  if s[0] != '0' && memo[s[1:]] == 0 {
    f(s[1:])
  }
  
  if len(s) > 1 && (s[0] == '1' || (s[0] == '2' && (s[1] == '0' || s[1] == '1' || s[1] == '2' || s[1] == '3' || s[1] == '4' || s[1] == '5' || s[1] == '6')) && memo[s[2:]] == 0) {
    f(s[2:])
  }
  
  memo[s] = 1

}
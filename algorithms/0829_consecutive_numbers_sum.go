import "math"

func consecutiveNumbersSum(N int) int {
  
  count := 1
  e := int(math.Sqrt(float64(N)))
  
  for i := 2; i <= e; i++ {
    if isExist(i, N) {
      count++
    }
  }
  
  return count
}

func isExist(L int, N int) bool {
  a := (2*N) % L
  b := ((2*N)/L) - L + 1
  return a == 0 && b > 0 && b % 2 == 0
}
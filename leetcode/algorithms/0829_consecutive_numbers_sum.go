import "math"

func consecutiveNumbersSum(N int) int {
  
  count := 0
  upper_limit := int(math.Sqrt(float64(2 * N) + 0.25) - 0.5)
  for k := 1; k <= upper_limit; k++ {
    if (N-k * (k+1)/2) % k == 0 {
      count++
    }
  }
  return count
}
var memo map[string]struct{}

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	ans := make([][]int, 0)
    memo = make(map[string]struct{}, 0)
	backtrace(nums, nil, &ans)
	return ans       
}

func backtrace(nums []int, prev []int, ans *[][]int) {
	if len(nums) == 0 {
        result := append([]int{}, prev...)
        key := fmt.Sprint(result)
        if _, ok := memo[key]; !ok {
            memo[key] = struct{}{}
    		*ans = append(*ans, result)
        }
		return
	}

	for i := 0; i < len(nums); i++ {
		backtrace(append(append([]int{}, nums[:i]...), nums[i+1:]...), append(prev, nums[i]), ans)
	}
}

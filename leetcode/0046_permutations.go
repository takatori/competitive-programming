func permute(nums []int) [][]int {
    
	if len(nums) == 0 {
		return nil
	}
	ans := make([][]int, 0)
	backtrace(nums, nil, &ans)
	return ans
}

func backtrace(nums []int, prev []int, ans *[][]int) {
	if len(nums) == 0 {
		*ans = append(*ans, append([]int{}, prev...))
		return
	}

	for i := 0; i < len(nums); i++ {
		backtrace(append(append([]int{}, nums[:i]...), nums[i+1:]...), append(prev, nums[i]), ans)
	}
}

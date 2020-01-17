func searchInsert(nums []int, target int) int {
    var i int
    for ; i < len(nums); i++ {
        if nums[i] == target {
            return i
        } else if nums[i] > target {
            return i
        } 
    }
    return i
}

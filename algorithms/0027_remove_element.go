func removeElement(nums []int, val int) int {
    var i int
    n := len(nums)    
    for i < n {
        if nums[i] == val {
            nums[i] = nums[n-1]
            n--
        } else {
            i++
        }
    }
    return n
}

func removeElement(nums []int, val int) int {
    var i int
    for i < len(nums) { // i = 3, len = 5, [0,1,2,3,4]
        j := i // j = 3
        for j < len(nums) - 1 && nums[j] == nums[j+1] { 
            j++ // j = 3
        }
        nums = append(nums[:i], nums[j:]...) // [0,1,2], [3,4] => [0,1,2,3,4]
        i = i + 1 // i = 4
    }
    return len(nums)
}


// [0,0,1,2,2,2,3,4]
// [0,1,2,2,2,3,4]
// [0,1,2,3,4]

func sortColors(nums []int)  {
    
    red, white, blue := 0, 0, len(nums)-1

    for white <= blue {
        if nums[white] == 0 {
            nums[red], nums[white] = nums[white], nums[red]
            red++
            white++
        } else if nums[white] == 1 {
            white++
        } else {
            nums[white], nums[blue] = nums[blue], nums[white]
            blue--
        }
    }
}

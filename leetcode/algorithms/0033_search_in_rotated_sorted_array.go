func search(nums []int, target int) int {
    
	n := len(nums)
	lo := 0
	hi := n - 1

	for lo < hi {
		mid := (lo+hi) / 2
		if nums[mid] > nums[hi] {
			lo = mid+1
		} else {
			hi = mid
		}
	}

	rot := lo
	lo = 0
	hi = n-1

	for lo <= hi {
		mid := (lo+hi) / 2
		realmid := (mid+rot) % n
		if nums[realmid] == target {
			return realmid
		} else if nums[realmid] < target {
			lo = mid+1
		} else {
			hi = mid-1
		}
	}

	return -1
}

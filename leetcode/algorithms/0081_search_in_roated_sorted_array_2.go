func search(nums []int, target int) bool {
    
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

  if nums[lo] >= target && target <= nums[n-1] {
    lo = hi
    hi = n-1
  } else {
    hi = lo
    lo = 0
  }

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
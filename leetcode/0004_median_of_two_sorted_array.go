func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
 
    i := 0
    j := 0
    nums1Len := len(nums1)
    nums2Len := len(nums2)
    n := nums1Len + nums2Len
    array := make([]int, 0, n)
    
    for i < nums1Len && j < nums2Len {
        if nums1[i] < nums2[j] {
            array = append(array, nums1[i])
            i++
        } else {
            array = append(array, nums2[j])
            j++
        }
    }
    
    if i == nums1Len {
        array = append(array, nums2[j:]...)
    } else {
        array = append(array, nums1[i:]...)
    }
    
    if n % 2 == 0 {
        return (float64(array[n/2-1]) + float64(array[n/2])) / 2
    } else {
        return float64(array[n/2])
    }
    
}

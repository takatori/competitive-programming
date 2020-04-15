/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func sortedArrayToBST(nums []int) *TreeNode {
  
	if len(nums) == 0 {
	  return nil
	}
	
	if len(nums) == 1 {
	  return &TreeNode {
		Val: nums[0],
		Left: nil,
		Right: nil,
	  }
	}
	  
	l := len(nums)
	
	v := nums[l/2]
	left := sortedArrayToBST(nums[:l/2])
	right := sortedArrayToBST(nums[l/2+1:])
	
	return &TreeNode{
	  Val: v,
	  Left: left,
	  Right: right,
	}
	
  }
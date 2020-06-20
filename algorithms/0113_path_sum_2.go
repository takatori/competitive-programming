/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func pathSum(root *TreeNode, sum int) [][]int {
    
	if root == nil {
	  return nil
	}
	
	if root.Left == nil && root.Right == nil {
	  if root.Val == sum {
		return [][]int{{root.Val}}
	  } else {
		return nil
	  }
	}
	
	s := sum - root.Val
	
	leftPath := pathSum(root.Left, s)
	rightPath := pathSum(root.Right, s)
	
	var result [][]int
	
	if leftPath != nil {
	  for _, l := range leftPath {
		result = append(result, append([]int{root.Val}, l...))
	  }
	}
	
	if rightPath != nil {
	  for _, r := range rightPath {
		result = append(result, append([]int{root.Val}, r...))
	  }
	}
	
	return result
  }
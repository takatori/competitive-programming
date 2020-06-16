/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func hasPathSum(root *TreeNode, sum int) bool {
    
	if root == nil {
	  return false
	}
	
	if root.Left == nil && root.Right == nil {
	  return root.Val == sum
	}
	
	n := sum - root.Val
	
	return hasPathSum(root.Left, n) || hasPathSum(root.Right, n)
  }
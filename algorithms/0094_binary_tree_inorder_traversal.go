/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func inorderTraversal(root *TreeNode) []int {

	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	curr := root
	
	for curr != nil || len(stack) > 0 {
	  for curr != nil {
		stack = append(stack, curr)
		curr = curr.Left
	  }
	  l := len(stack)-1
	  curr = stack[l]
	  stack = stack[:l]
	  res = append(res, curr.Val)
	  curr = curr.Right
	}
	
	return res
	
	
  }
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func flatten(root *TreeNode)  {
    
	if root == nil {
	  return
	}
	
	if root.Left == nil && root.Right == nil {
	  return
	}
	
	flatten(root.Left)
	flatten(root.Right)
	
	if root.Left != nil {
	  r := root.Right
	  root.Right = root.Left
	  root.Left = nil
	  t := root.Right
	  for t.Right != nil {
		t = t.Right
	  }
	  t.Right = r
	}
	
  }
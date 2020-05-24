
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isBalanced(root *TreeNode) bool {
    
	if root == nil {
	  return true
	}
	
	lh, rh := hight(root.Left), hight(root.Right)
	
	if lh > rh {
	  return (lh - rh) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
	} else {
	  return (rh - lh) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
	}
   
  } 
  
  func hight(node *TreeNode) int {	

	if node == nil {
	  return 0
	}
	
	if node.Left == nil && node.Right == nil {
	  return 1
	}
	
	var lh, rh int
	
	if node.Left != nil {
	  lh = hight(node.Left)
	}
	
	if node.Right != nil {
	  rh = hight(node.Right)
	}
	
	if lh > rh {
	  return lh + 1
	} else {
	  return rh + 1
	}
  }
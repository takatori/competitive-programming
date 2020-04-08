/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func levelOrderBottom(root *TreeNode) [][]int {
    
	if root == nil {
	  return [][]int{}
	}
	
	result := make([][]int, 0)
	q := []*TreeNode{root}
	
	for len(q) > 0 {
	  
	  tmp := make([]int, 0)
	  next := make([]*TreeNode, 0)
	  
	  for i := 0; i < len(q); i++ {
		tmp = append(tmp, q[i].Val)    
		if q[i].Left != nil {
		  next = append(next, q[i].Left)
		}
		if q[i].Right != nil {
		  next = append(next, q[i].Right)
		}
	  }
	  
	  result = append([][]int{tmp}, result...)
	  q = next
	}
	
	return result    
  }
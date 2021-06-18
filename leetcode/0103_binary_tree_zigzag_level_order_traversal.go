/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {

  if root == nil {
    return [][]int{}
  }
  
  result := [][]int{}
  
  q := []*TreeNode{root}
  
  for len(q) > 0 {
    next := make([]*TreeNode, 0)
    tmp := []int{}
    for _, v := range q {
      if len(result) % 2 == 0 {
        tmp = append(tmp, v.Val)
      } else {
        tmp = append([]int{v.Val}, tmp...)
      }
      
      if v.Left != nil {
        next = append(next, v.Left)
      }
      
      if v.Right != nil {
        next = append(next, v.Right)
      }
    }
    q = next
    result = append(result, tmp)
  }
  
  return result
}
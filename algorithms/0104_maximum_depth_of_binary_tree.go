/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    
    if root == nil {
        return 0
    }
    
    var result int
    q := []*TreeNode{root}
    
    for len(q) > 0 {
        c := []*TreeNode{}
        for i := 0; i < len(q); i++ {
            if q[i] == nil {
                continue
            }
            if q[i].Left != nil {
                c = append(c, q[i].Left)
            }
            if q[i].Right != nil {
                c = append(c, q[i].Right)
            }
        }
        q = c
        result++
    }
    
    return result
}

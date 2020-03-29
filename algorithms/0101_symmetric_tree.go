/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    
    if root == nil || (root.Left == nil && root.Right == nil) {
        return true
    } else if root.Left == nil || root.Right == nil {
        return false
    }
    
    queue := make([]*TreeNode, 0)
    
    for len(queue) > 0 {
        
        nextQueue := make([]*TreeNode, 0)
        l := 0
        r := len(queue)-1
        
        for l < r {
            if (queue[l].Value != queue[r].Value) || (queue[l].Left == nil && queue[r].Right != nil) || {
                return false
            }
            if queue[l].Right != nil {
                nextQueue = append([]*Tree{queue[l].Right}, append(nextQueue, queue[r].Left)...)
            }
            if queue[l].Left != nil {
                 nextQueue = append([]*Tree{queue[l].Left}, append(nextQueue, queue[r].Right)...)
            }
            l++
            r--
        }
    }

    return true
}


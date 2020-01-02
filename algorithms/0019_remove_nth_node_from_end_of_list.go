/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    
    nodes := make([]*ListNode, 0)
    var length int
    
    if head == nil {
        return nil
    }
    
    for n := head; n != nil; n = n.Next {
        nodes = append(nodes, n)
        length++
    }
    
    if length == n {
          head = head.Next
    } else if n == 1 {
          nodes[length-2].Next = nil
    } else {
        nodes[length-n-1].Next = nodes[length-n+1]
    }
    
   return head
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    
    prev := head
    current := head
    
    for ; current != nil; current = current.Next {
        if prev.Val != current.Val {
            prev.Next = current
            prev = current
        }
        if current.Next == nil {
            prev.Next = nil
        }
    }
    
    return head    
}

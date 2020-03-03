/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    
    if head == nil {
        return head
    }
    
    var length int
    var last *ListNode
    
    for h := head; h != nil; h = h.Next {
        length++
        last = h
    }
    
    p := k % length
    
    if p == 0 {
        return head
    }
    
    last.Next = head
    
    var prev *ListNode
   
    for i := 0; i < length - p; i++ {
        prev = head
        head = head.Next
    }
    
    prev.Next = nil
    
    return head
    
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    
    if head == nil || head.Next == nil {
        return head
    } 
    var prev *ListNode
    top := head.Next
    current := head 
    for current != nil { 
        swap(prev, current, current.Next)
        prev = current
        current = current.Next
    }
    return top
}

func swap(a *ListNode, b *ListNode, c *ListNode) {
    if a == nil {
        b.Next = c.Next 
        c.Next = b
        return
    }
    if c == nil {
        return
    }
    a.Next = b.Next
    b.Next = c.Next
    c.Next = b
    return
}

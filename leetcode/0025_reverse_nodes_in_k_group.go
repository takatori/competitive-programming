/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    if k <= 1 {
        return head
    }
    nodesLen := 0
    for p := head; p != nil; p = p.Next {
        nodesLen++
    }
    
    rhead := new(ListNode)
    rhead.Next = head
    pre := rhead
    tail,next := head,head
    
    for ;nodesLen >= k; nodesLen -= k {
        for i := 0; i < k; i++ {
            temp := next.Next
            next.Next = pre.Next
            pre.Next = next
            next = temp
            tail.Next = next
        }
        pre = tail
        tail = next
    }
    
    return rhead.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    top := head // 1
    for i := k; i > 1; i-- { // k = 2
        top = top.Next
    }
    // top = 2
        
    var p *ListNode
    current := head 
    for current != nil { // 2->1->4->3->5, current = 5
        prev := p // 3
        start, end := current, current // 5, 5
        for i := k; i > 1 && start != nil; i-- { // i = 2 -> 1
            for j := i; j > 1 && end != nil; j-- { // j = 2
                end = end.Next // 5 -> nil
            }
            swapPare(prev, start, end) // (3, 5, nil) => 2->1->4->3->5
            prev = start // prev = 5
            start = start.Next // start = nil
            end = start // end = nil
        }
        p = current // 5
        current = current.Next // nil
    }
    
    return top
}

func swapPare(prev, current, next *ListNode) {
    if prev == nil {
        current.Next = next.Next 
        next.Next = current
        return
    }
    if next == nil {
        return
    }
    prev.Next = current.Next
    current.Next = next.Next
    next.Next = current
}

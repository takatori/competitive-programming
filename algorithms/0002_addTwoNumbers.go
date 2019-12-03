func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    
    var head *ListNode
    var tail *ListNode
    var carry int
    
    for l1 != nil || l2 != nil || carry != 0 {
        
        var l1Val int
        var l2Val int
        
        if l1 != nil {
            l1Val = l1.Val
            l1 = l1.Next
        }
        
        if l2 != nil {
            l2Val = l2.Val
            l2 = l2.Next
        }
            
        sum := l1Val + l2Val + carry
        carry = sum / 10
        node := &ListNode{Val: sum % 10, Next: nil}
        
        if head == nil {           
            head = node
            tail = node
        } else {
            tail.Next = node
            tail = node
        }
    }
    if (carry > 0) {
        tail.Next = &ListNode{Val: carry, Next: nil}
    }
    
    return head
}

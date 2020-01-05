func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    
    head := &ListNode{}
    last := head
    
    for l1 != nil && l2 != nil {
        if l1.Val <= l2.Val {
            last.Next = l1
            last = last.Next
            l1 = l1.Next
        } else {
            last.Next = l2
            last = last.Next
            l2 = l2.Next
        }
    }
    
    if l1 == nil {
        last.Next = l2
    } else {
        last.Next = l1
    }
    
    return head.Next
}

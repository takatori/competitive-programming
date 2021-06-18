/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
	  return head
	}
	
	c, n := head, head.Next
	var result *ListNode
	var ans *ListNode
	
	for c != nil {
	  count := 0
	  for n != nil && c.Val == n.Val {
		n = n.Next
		count++
	  }
	  if count == 0 {
		if result == nil {
		  ans = &ListNode{
			Val: c.Val,
			Next: nil,
		  }
		  result = ans
		} else {
		  result.Next = c
		  result = c
		}
	  }
	  c = n
	  if n != nil {
		n = n.Next
	  }
	}
	
	if result != nil {
	  result.Next = nil
	}
	
	return ans
	
  }
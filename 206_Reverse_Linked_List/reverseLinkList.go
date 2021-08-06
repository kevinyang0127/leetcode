package reverseLinkList

type ListNode struct {
	Val  int
	Next *ListNode
}

//Iterative
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prevNode *ListNode
	currNode := head
	nextNode := head.Next

	for nextNode != nil {
		currNode.Next = prevNode
		prevNode = currNode
		currNode = nextNode
		nextNode = nextNode.Next
	}
	currNode.Next = prevNode
	return currNode
}

//Iterative, create new list
func reverseList_2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var reverseHead *ListNode
	ptr := head
	for ptr.Next != nil {
		reverseHead = addValToHead(reverseHead, ptr.Val)
		ptr = ptr.Next
	}
	reverseHead = addValToHead(reverseHead, ptr.Val)
	return reverseHead
}

func addValToHead(listHead *ListNode, val int) *ListNode {
	newNode := &ListNode{
		Val:  val,
		Next: listHead,
	}

	return newNode
}

//recursive
func reverseList_recursive(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	reverseHead := reverseList_recursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return reverseHead
}

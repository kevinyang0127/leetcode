package reverseLinkList

import (
	"testing"
)

func createLinkList(vals []int) *ListNode {
	var prevNode *ListNode
	var head *ListNode
	for _, v := range vals {
		newNode := &ListNode{
			Val: v,
		}

		if prevNode != nil {
			prevNode.Next = newNode
		} else {
			head = newNode
		}

		prevNode = newNode
	}
	return head
}

func linkListToSlice(head *ListNode) []int {
	ptr := head
	s := []int{}
	for ptr != nil {
		s = append(s, ptr.Val)
		ptr = ptr.Next
	}
	return s
}

func isSliceEq(s1 []int, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func TestReverseList(t *testing.T) {
	head := createLinkList([]int{1, 2, 3, 4, 5})
	reverse := reverseList(head)
	ans := []int{5, 4, 3, 2, 1}

	if isSliceEq(linkListToSlice(reverse), ans) {
		t.Log("Accept")
	} else {
		t.Error("wrong ans")
	}
}

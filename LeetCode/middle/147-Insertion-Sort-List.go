package main

import (
	"ark"
)

func insertionSortList(head *ark.ListNode) *ark.ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ark.ListNode{Next: head}
	pointerA, pointerB := head, head.Next
	for pointerB != nil {
		if pointerA.Val <= pointerB.Val {
			pointerA = pointerA.Next
		} else {
			prev := dummyHead
			for head.Next.Val <= pointerB.Val {
				prev = prev.Next
			}
			pointerA.Next = pointerB.Next
			pointerB.Next, prev.Next = prev.Next, pointerB
		}
		pointerB = pointerA.Next
	}
	return dummyHead.Next
}

func main() {

}

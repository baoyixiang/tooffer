package main

type ListNode struct {
	val int
	next *ListNode
}

func NewLinkedList(arr []int) *ListNode {
	head := &ListNode{0, nil}

	if len(arr) == 0 {
		return head
	}

	tmpPoint := head
	for _, item := range arr {
		curNode := &ListNode{item, nil}
		tmpPoint.next = curNode
		tmpPoint = tmpPoint.next
	}

	return head
}

func ()  {
	
}
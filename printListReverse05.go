package main

import (
	"container/list"
	"fmt"
)

// 反转打印链表有三种方法
// 1. 反转原链表， 顺序打印
// 2. 利用栈
// 3. 递归
func printListReverse(head *ListNode) {
	if head == nil {
		return
	}
	stack := list.List{}

	tmpPointer := head
	for tmpPointer.next != nil {
		stack.PushBack(tmpPointer.next.val)
		tmpPointer = tmpPointer.next
	}

	for stack.Len() > 0 {
		curNode := stack.Back()
		fmt.Println(curNode.Value)
		stack.Remove(curNode)
	}
}
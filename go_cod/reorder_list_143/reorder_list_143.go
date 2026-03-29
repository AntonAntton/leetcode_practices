package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// Step 1: Find the middle of the list
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Step 2: Reverse the second half
	secondHalf := slow.Next
	slow.Next = nil
	var prev *ListNode
	for secondHalf != nil {
		temp := secondHalf.Next
		secondHalf.Next = prev
		prev = secondHalf
		secondHalf = temp
	}

	// Step 3: Merge the two halves
	firstHalf := head
	for prev != nil {
		temp1 := firstHalf.Next
		temp2 := prev.Next
		firstHalf.Next = prev
		prev.Next = temp1
		firstHalf = temp1
		prev = temp2
	}
}

func main() {
	// Example usage:
	head := &ListNode{Val: 3}
	head.Next = &ListNode{Val: 3}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	reorderList(head)

	// Print the reordered list
	current := head
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
}

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	first := head
	second := head.Next
	nextPairHead := second.Next
	second.Next = first
	first.Next = swapPairs(nextPairHead)
	return second
}

func main() {
	// Example usage:
	// Create a linked list: 1 -> 2 -> 3 -> 4
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}

	// Swap pairs
	newHead := swapPairs(head)

	// Print the swapped linked list
	current := newHead
	for current != nil {
		println(current.Val)
		current = current.Next
	}
}

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func solve(head ListNode) {
	slow, fast := head, head

	for fast.Next != nil {
		slow = *slow.Next
		fast = *fast.Next.Next
	}

	curr := &slow
	var prev *ListNode
	while

}

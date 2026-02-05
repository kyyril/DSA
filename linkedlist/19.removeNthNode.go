func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{0, head}
    fast, slow := dummy, dummy

    // Move fast n+1 steps
    for i := 0; i <= n; i++ {
        fast = fast.Next
    }

    // Move fast & slow together
    for fast != nil {
        fast = fast.Next
        slow = slow.Next
    }

    // Remove nth node
    slow.Next = slow.Next.Next

    return dummy.Next
}

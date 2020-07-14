package linked_lists

func ReverseLinkedList(head *LinkedList) *LinkedList {
	var p1, p2 *LinkedList = nil, head
	for p2 != nil {
		p3 := p2.Next
		p2.Next = p1
		p1 = p2
		p2 = p3
	}
	return p1
}

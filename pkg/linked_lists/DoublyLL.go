package linked_lists

type Node struct {
	Value      int
	Prev, Next *Node
}

type DoublyLinkedList struct {
	Head, Tail *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (ll *DoublyLinkedList) SetHead(node *Node) {
	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
		return
	}
	ll.InsertBefore(ll.Head, node)
}

func (ll *DoublyLinkedList) SetTail(node *Node) {
	if ll.Tail == nil {
		ll.Head = node
		ll.Tail = node
		return
	}
	ll.InsertAfter(ll.Tail, node)
}

func (ll *DoublyLinkedList) InsertBefore(node, nodeToInsert *Node) {
	if nodeToInsert == ll.Head && nodeToInsert == ll.Tail {
		return
	}
	ll.Remove(nodeToInsert)
	nodeToInsert.Next = node
	nodeToInsert.Prev = node.Prev
	if node.Prev == nil {
		ll.Head = nodeToInsert
	} else {
		node.Prev.Next = nodeToInsert
	}
	node.Prev = nodeToInsert
}

func (ll *DoublyLinkedList) InsertAfter(node, nodeToInsert *Node) {
	if nodeToInsert == ll.Head && nodeToInsert == ll.Tail {
		return
	}
	ll.Remove(nodeToInsert)
	nodeToInsert.Prev = node
	nodeToInsert.Next = node.Next
	if node.Next == nil {
		ll.Tail = nodeToInsert
	} else {
		node.Next.Prev = nodeToInsert
	}
	node.Next = nodeToInsert
}

func (ll *DoublyLinkedList) InsertAtPosition(position int, nodeToInsert *Node) {
	if position == 1 {
		ll.SetHead(nodeToInsert)
		return
	}
	node := ll.Head
	curr := 1
	for node != nil && curr != position {
		node = node.Next
		curr++
	}
	if node != nil {
		ll.InsertBefore(node, nodeToInsert)
	} else {
		ll.SetTail(nodeToInsert)
	}
}

func (ll *DoublyLinkedList) RemoveNodesWithValue(value int) {
	node := ll.Head
	for node != nil {
		curr := node
		node = node.Next
		if curr.Value == value {
			ll.Remove(curr)
		}
	}
}

func (ll *DoublyLinkedList) Remove(node *Node) {
	if node == ll.Head {
		ll.Head = ll.Head.Next
	}
	if node == ll.Tail {
		ll.Tail = ll.Tail.Prev
	}
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	node.Prev = nil
	node.Next = nil
}

func (ll *DoublyLinkedList) ContainsNodeWithValue(value int) bool {
	node := ll.Head
	for node != nil && node.Value != value {
		node = node.Next

	}
	return node != nil
}

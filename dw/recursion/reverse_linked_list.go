package recursion

import . "basics/dw/ds"

func ReverseLinkedList(list *LinkedList) *LinkedList {
	head := list.Head
	doReverse(head, nil)
	return list
}

func doReverse(curr *Node, prev *Node) *Node {
	if curr.Next == nil {
		curr.Next = prev
		return curr
	}
	doReverse(curr.Next, curr)
	curr.Next = prev
	return curr
}

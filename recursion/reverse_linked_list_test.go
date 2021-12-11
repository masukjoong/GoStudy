package recursion

import "testing"
import . "basics/ds"

func TestReverseLinkedList(t *testing.T) {
	tail := &Node{Data: 4}
	head := &Node{Data: 1, Next: &Node{Data: 2, Next: &Node{Data: 3, Next: tail}}}
	linkedList := &LinkedList{Head: head, Tail: tail}
	ReverseLinkedList(linkedList)
	curr := tail
	for i:=4;i>0;i-- {
		if curr.Data != i {
			t.Errorf("curr.Data = %d; want %d", curr.Data, i)
		}
		curr = curr.Next
	}
}

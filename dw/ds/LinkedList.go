package ds

type LinkedList struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Data int
	Next *Node
}

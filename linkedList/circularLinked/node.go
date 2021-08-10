package circularLinked

//	the doubly circular linked list does not contain the NULL value in the previous field of the node.
type Node struct {
	PreviousNode *Node
	NextNode *Node
	Value	string 		// could be any type
}

type CircularLinkedList struct {
	FirstNode *Node
}
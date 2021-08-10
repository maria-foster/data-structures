package circularLinked

//	the doubly circular linked list does not contain the NULL value in the previous field of the node.
type Node struct {
	previousNode *Node
	nextNode *Node
	value	string 		// could be any type
}

type CircularLinkedList struct {
	firstNode *Node
}
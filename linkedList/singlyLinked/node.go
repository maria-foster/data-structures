package singlyLinked

type Node struct {
	nextNode *Node
	value	string 		// could be any type
}

type SinglyLinkedList struct {
	firstNode *Node
}


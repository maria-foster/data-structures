package doublyLinked

type Node struct {
	previousNode *Node
	nextNode *Node
	value	string 		// could be any type
}

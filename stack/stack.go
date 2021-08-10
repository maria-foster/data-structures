package stack

import (
	lists "dataStructures/linkedList/circularLinked"
)
type Stack struct {
	List	lists.CircularLinkedList
}
func (stack Stack) Pop() (*lists.Node) {
	//  inserting and removing elements from the end of the list
	// firstNode := stack.List.FirstNode
	// poppedNode := stack.List.FirstNode.PreviousNode
	// nextNode := stack.List.FirstNode.PreviousNode.NextNode
	stack.List.FirstNode.PreviousNode = stack.List.FirstNode.PreviousNode.NextNode
	stack.List.FirstNode.PreviousNode.NextNode.NextNode = stack.List.FirstNode
	return stack.List.FirstNode.PreviousNode

}
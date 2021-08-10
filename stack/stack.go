package stack

import (
	lists "dataStructures/linkedList/circularLinked"
)
type Stack struct {
	List	lists.CircularLinkedList
}
// FOUR FUNCTHIONS: pop, push, peek, isEmpty

func (stack Stack) Pop() (*lists.Node) {
	//  inserting and removing elements from the end of the list
	// firstNode := stack.List.FirstNode
	// poppedNode := stack.List.FirstNode.PreviousNode
	// nextNode := stack.List.FirstNode.PreviousNode.NextNode
	stack.List.FirstNode.PreviousNode = stack.List.FirstNode.PreviousNode.NextNode
	stack.List.FirstNode.PreviousNode.NextNode.NextNode = stack.List.FirstNode
	return stack.List.FirstNode.PreviousNode

}

func (stack Stack) Push( newNode *lists.Node) (*lists.Node) {
	//  inserting and removing elements from the end of the list
	// firstNode := stack.List.FirstNode
	// pushedNode := newNode
	// nextNode := stack.List.FirstNode.PreviousNode.NextNode
	newNode.NextNode = stack.List.FirstNode
	newNode.PreviousNode =  stack.List.FirstNode.PreviousNode.NextNode
	stack.List.FirstNode.PreviousNode = newNode
	stack.List.FirstNode.PreviousNode.NextNode = newNode
	return stack.List.FirstNode.PreviousNode

}

func (stack Stack) IsEmpty() (bool){
	return stack.List.FirstNode == nil
}

func (stack Stack) Peek() (*lists.Node) {
	// returns the node at the top of the stack
	return stack.List.FirstNode.PreviousNode
}
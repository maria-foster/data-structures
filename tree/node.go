package tree


type Node struct {
	parentNode 		*Node
	childrenNodes 	[]*Node
	value	string 		// could be any type
}
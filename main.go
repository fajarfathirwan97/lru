package list

type Node struct {
	Data string `json:"data"`
	prev *Node
	next *Node
}

func (n *Node) Next() *Node {
	return n.next
}
func (n *Node) Prev() *Node {
	return n.prev
}

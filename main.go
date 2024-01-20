package list

type Node struct {
	Data string `json:"data"`
	prev *Node
	next *Node
}

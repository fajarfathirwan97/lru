package list

import "log"

type DoublyLinkedList struct {
	len  int64
	Head *Node `json:"Head"`
	Tail *Node `json:"Tail"`
}

func (d *DoublyLinkedList) Append(data string) *Node {
	if d.Head == nil && d.Tail == nil {
		return d.Add(data)
	}
	newNode := &Node{Data: data, prev: d.Tail}
	d.Tail.next = newNode
	d.Tail = newNode
	d.len++
	return newNode
}
func (d *DoublyLinkedList) Prepend(data string) *Node {
	if d.Head == nil && d.Tail == nil {
		return d.Add(data)
	}
	newNode := &Node{Data: data, next: d.Head}
	d.Head.prev = newNode
	d.Head = newNode
	d.len++
	return newNode
}
func (d *DoublyLinkedList) Add(data string) *Node {
	d.len++
	if d.Head == nil && d.Tail == nil {
		d.Head = &Node{Data: data}
		d.Tail = d.Head
		return d.Tail
	} else {
		return d.Prepend(data)
	}
}

func (d *DoublyLinkedList) RemoveAtNode(node *Node) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		d.RemoveHead()
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		d.RemoveTail()
	}
	node.next = nil
	node.prev = nil
}
func (d *DoublyLinkedList) Remove(data string) {
	if d.Head == nil || d.Tail == nil {
		return
	}
	if d.Tail.Data == data {
		d.RemoveTail()
		return
	}
	if d.Head.Data == data {
		d.RemoveHead()
		return
	}
	node := d.SearchNode(d.Head, data)
	d.RemoveAtNode(node)
}

func (d *DoublyLinkedList) RemoveTail() {
	d.Tail = d.Tail.prev
	d.Tail.next = nil
	d.len--
}

func (d *DoublyLinkedList) RemoveHead() {
	d.Head = d.Head.next
	d.Head.prev = nil
	d.len--
}
func (d *DoublyLinkedList) Update(old, new string) (*Node, bool) {
	found := d.SearchNode(d.Head, old)
	if found != nil {
		found.Data = new
		return found, true
	}
	return nil, false
}
func (d *DoublyLinkedList) InsertAtNode(node *Node, data string) *Node {
	newNode := &Node{Data: data, next: node.next, prev: node}
	node.next.prev = newNode
	node.next = newNode
	return newNode
}
func (d *DoublyLinkedList) InsertAtIndex(index int64, data string) *Node {
	if index == 0 {
		return d.Prepend(data)
	}
	if index == d.len {
		return d.Append(data)
	}
	node := d.Head
	for i := int64(0); i < index-1; i++ {
		node = node.next
	}
	newNode := d.InsertAtNode(node, data)
	d.len++
	return newNode
}
func (d *DoublyLinkedList) SearchNode(node *Node, data string) *Node {
	for node != nil {
		if node.Data == data {
			return node
		}
		return d.SearchNode(node.next, data)
	}
	return nil
}
func (d *DoublyLinkedList) Search(node *Node, data string) string {
	for node != nil {
		if node.Data == data {
			return node.Data
		}
		return d.Search(node.next, data)
	}
	return ""
}
func (d *DoublyLinkedList) Debug() {
	nextNode := d.Head
	for nextNode != nil {
		log.Println(nextNode)
		nextNode = nextNode.next
	}
	log.Println(d.len, "LENGT")
	log.Println("============")
}
func (d *DoublyLinkedList) Len() int64 { return d.len }

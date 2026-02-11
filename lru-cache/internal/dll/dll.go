package dll

import "sync"

type DLL struct {
	Head     *Node
	Tail     *Node
	capacity int
	mu       sync.RWMutex
}

func NewDLL(cap int) *DLL {
	head := &Node{}
	tail := &Node{}

	head.Next = tail
	tail.Prev = head

	return &DLL{
		Head:     head,
		Tail:     tail,
		capacity: cap,
	}
}

func (d *DLL) AddFirst(new_node *Node) {
	first := d.Head.Next
	d.Head.Next = new_node
	new_node.Next = first
	new_node.Prev = d.Head
	first.Prev = new_node
}

func (d *DLL) RemoveNode(n *Node) {
	prev := n.Prev
	next := n.Next
	prev.Next = next
	next.Prev = prev
}
func (d *DLL) RemoveBeginning() *Node {
	if d.Head.Next == d.Tail {
		return nil
	}
	node := d.Head.Next
	next_node := node.Next
	d.Head.Next = next_node
	next_node.Prev = d.Head
	return node
}

func (d *DLL) InsertAtLast(new_node *Node) {
	last := d.Tail.Prev
	last.Next = new_node
	d.Tail.Prev = new_node
	new_node.Next = d.Tail
	new_node.Prev = last
}

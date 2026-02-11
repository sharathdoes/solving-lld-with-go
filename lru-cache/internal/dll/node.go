package dll

type Node struct {
	Key string
	Value string
	Prev *Node
	Next *Node
}

func NewNode(key, value string) *Node {
	return &Node{
		Key: key, Value: value,Prev: nil, Next: nil,
	}
}
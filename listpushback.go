package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	new := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = new
	} else {
		new.Next = l.Head
		l.Head = new
	}
}

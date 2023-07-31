package piscine

func ListReverse(l *List) {
	act := l.Head
	var prev *NodeL
	for act != nil {
		next := act.Next
		act.Next = prev
		prev = act
		act = next
	}
	l.Head = prev
}

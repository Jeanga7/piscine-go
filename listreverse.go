package piscine

func ListReverse(l *List) {
	debut := l.Head
	var act *NodeL
	for debut != nil {
		next := debut
		debut = act
		act = act.Next
		act.Next = next
	}
	l.Head = act
}

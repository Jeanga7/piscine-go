package piscine

func ListReverse(l *List) {
	deb, actuel, prev := l.Head, l.Head, l.Head
	//prev = nil

	for actuel != nil {
		next := actuel.Next
		actuel.Next = prev
		prev = actuel
		actuel = next
	}
	l.Head = prev
	l.Tail = deb
	// l.Tail.Next = nil
}

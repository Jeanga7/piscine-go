package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	count := 0
	for l != nil {
		count++
		l = l.Next
		if count == pos {
			return l
		}
	}
	return nil
}

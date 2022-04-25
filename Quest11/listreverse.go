package piscine

func ListReverse(l *list) {
	current := l.Head
	previous := l.Head
	previous = nil
	for current != nil {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}
	l.Head = previous
}

package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	n := &NodeI{}
	n.Data = data_ref
	n.Next = nil

	if l == nil || l.Data >= n.Data {
		n.Next = l
		return n
	} else {
		tmpl := l
		for tmpl.Next != nil && tmpl.Next.Data < n.Data {
			tmpl = tmpl.Next
		}
		n.Next = tmpl.Next
		tmpl.Next = n
	}
	return l
}

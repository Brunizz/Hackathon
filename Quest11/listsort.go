package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	Head := l
	if Head == nil {
		return nil
	}
	Head.Next = ListSort(Head.Next)

	if Head.Next != nil && Head.Data > Head.Next.Data {
		Head = replace(Head)
	}
	return Head
}

func replace(l *NodeI) *NodeI {
	a := l
	b := l.Next
	cantusereturnhere := b
	for n != nil && l.Data > b.Data {
		a = b
		b = b.Next
	}
	a.Next = l
	l.Next = b
	return cantusereturnhere
}

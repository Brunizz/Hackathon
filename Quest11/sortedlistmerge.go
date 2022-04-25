package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func SortedListMerge(l1 *NodeI, l2 *NodeI) *NodeI {

	l1 = ListSort(l1)
	l2 = ListSort(l2)

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Data <= l2.Data {
		l1.Next = SortedListMerge(l1.Next, l2)
		return l1

	} else {

		l2.Next = SortedListMerge(l1, l2.Next)
		return l2
	}
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

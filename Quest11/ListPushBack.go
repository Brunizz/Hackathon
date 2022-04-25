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
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
	} else {
		listInsert := l.Head
		for listInsert.Next != nil {
			listInsert = listInsert.Next
		}
		listInsert.Next = n
		l.Tail = n
	}
}

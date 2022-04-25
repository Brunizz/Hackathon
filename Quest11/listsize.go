package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	j := l.Head
	i := 0
	for j != nil {
		i++
		j = j.Next
	}
	return i
}

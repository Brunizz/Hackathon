package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil {
		return nil
	}
	counter := 1
	index := 0
	for index < pos {
		if counter.Next == nil {
			return nil
		}
		counter = counter.Next
		index++
	}
	return counter
}

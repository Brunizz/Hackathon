package main

import "fmt"

type node struct { // Description of the node, needs data and a pointer to the next node. In this case we will use a single number, therefore "data int" is enough.
	data int
	next *node
}

type linkedList struct { // The actual list, needs to hold a head with an address pointing to the node. Length is for specifying the length of the list.
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) { // Sets linkedList as a receiver, prepend will take a node to be added at the front
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printListData() { // Used to print out the data in the nodes
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0 { // if length is zero (empty list), immediately return to avoid runtime error
		return
	}
	if l.head.data == value { // if the value to be deleted is the header, make the second node the header
		l.head = l.head.next
		l.length--
		return
	}
	previousToDelete := l.head
	for previousToDelete.next.data != value { // Comparing data in the next node because it's harder to delete a node that we've run past, since we don't have any pointers to previous nodes address
		if previousToDelete.next.next == nil { // If you reach the end of the list without finding the value to delete, return to avoid runtime error
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	mylist := linkedList{}
	node1 := &node{data: 48}
	node2 := &node{data: 18}
	node3 := &node{data: 16}
	node4 := &node{data: 11}
	node5 := &node{data: 7}
	node6 := &node{data: 2}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)
	mylist.prepend(node6)
	fmt.Println(mylist)           // Prints the address of the head (node 6) and the length of the list
	mylist.printListData()        // Prints using func (l linkedList) printListData
	mylist.deleteWithValue(16)    // Deletes the value from a set node
	mylist.printListData()        // Prints using func (l linkedList) printListData
	mylist.deleteWithValue(100)   // Tries to delete a value that does not exist
	mylist.deleteWithValue(2)     // In this case, deletes the head from the list
	mylist.printListData()        // Prints using func (l linkedList) printListData
	emptyList := linkedList{}     // Sets a new empty list
	emptyList.deleteWithValue(10) // Tries to delete a value that does not exist from a list that is empty
}

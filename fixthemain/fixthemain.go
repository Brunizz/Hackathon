package main

import "fmt"

type Door struct {
	state bool
}

const OPEN = false
const CLOSE = true

func OpenDoor(checkDoor *Door) {
	fmt.Println("Door Opening...")
	checkDoor.state = OPEN
}

func CloseDoor(checkDoor *Door) {
	fmt.Println("Door Closing...")
	checkDoor.state = CLOSE
}

func IsDoorOpen(checkDoor Door) bool {
	fmt.Println("is the Door open ?")
	return checkDoor.state
}

func IsDoorClose(checkDoor Door) bool {
	fmt.Println("is the Door close ?")
	return checkDoor.state
}

func main() {
	door := Door{}

	OpenDoor(&door)
	if IsDoorClose(door) {
		OpenDoor(&door)
	}
	if IsDoorOpen(door) {
		CloseDoor(&door)
	}
	if door.state == OPEN {
		CloseDoor(&door)
	}
}

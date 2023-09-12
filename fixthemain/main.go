package main

import "github.com/01-edu/z01"

type Door struct {
	state bool
}

const (
	OPEN  = false
	CLOSE = true
)

func OpenDoor(ptrDoor *Door) {
	z01.PrintRune('D')
	z01.PrintRune('o')
	z01.PrintRune('o')
	z01.PrintRune('r')
	z01.PrintRune(' ')
	z01.PrintRune('O')
	z01.PrintRune('p')
	z01.PrintRune('e')
	z01.PrintRune('n')
	z01.PrintRune('i')
	z01.PrintRune('n')
	z01.PrintRune('g')
	z01.PrintRune('.')
	z01.PrintRune('.')
	z01.PrintRune('.')
	z01.PrintRune('\n')
	ptrDoor.state = OPEN
}

func CloseDoor(ptrDoor *Door) {
	z01.PrintRune('D')
	z01.PrintRune('o')
	z01.PrintRune('o')
	z01.PrintRune('r')
	z01.PrintRune(' ')
	z01.PrintRune('C')
	z01.PrintRune('l')
	z01.PrintRune('o')
	z01.PrintRune('s')
	z01.PrintRune('i')
	z01.PrintRune('n')
	z01.PrintRune('g')
	z01.PrintRune('.')
	z01.PrintRune('.')
	z01.PrintRune('.')
	z01.PrintRune('\n')
	ptrDoor.state = CLOSE
}

func IsDoorOpen(ptrDoor Door) bool {
	z01.PrintRune('i')
	z01.PrintRune('s')
	z01.PrintRune(' ')
	z01.PrintRune('t')
	z01.PrintRune('h')
	z01.PrintRune('e')
	z01.PrintRune(' ')
	z01.PrintRune('D')
	z01.PrintRune('o')
	z01.PrintRune('o')
	z01.PrintRune('r')
	z01.PrintRune(' ')
	z01.PrintRune('o')
	z01.PrintRune('p')
	z01.PrintRune('e')
	z01.PrintRune('n')
	z01.PrintRune('e')
	z01.PrintRune('d')
	z01.PrintRune(' ')
	z01.PrintRune('?')
	z01.PrintRune('\n')
	return ptrDoor.state
}

func IsDoorClose(ptrDoor Door) bool {
	z01.PrintRune('i')
	z01.PrintRune('s')
	z01.PrintRune(' ')
	z01.PrintRune('t')
	z01.PrintRune('h')
	z01.PrintRune('e')
	z01.PrintRune(' ')
	z01.PrintRune('D')
	z01.PrintRune('o')
	z01.PrintRune('o')
	z01.PrintRune('r')
	z01.PrintRune(' ')
	z01.PrintRune('c')
	z01.PrintRune('l')
	z01.PrintRune('o')
	z01.PrintRune('s')
	z01.PrintRune('e')
	z01.PrintRune('d')
	z01.PrintRune(' ')
	z01.PrintRune('?')
	z01.PrintRune('\n')
	return ptrDoor.state
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

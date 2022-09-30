package main

import (
	"fmt"
)

func main() {
	fmt.Println("initializing board!")

	board := [Height][Width]Slot{}

	fmt.Println("Player 1 (Black) begins!")

	board[4][4] = Black

	PrettyPrint(board)
}

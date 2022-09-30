package main

import "fmt"

func PrettyPrint(board [Height][Width]Slot) {
	for _, arr := range board {
		fmt.Println(arr)
	}
}

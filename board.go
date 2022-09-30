package main

const Height = 9
const Width = 9

type Slot int64

const (
	Empty Slot = iota
	Black
	White
)

// type Board = [][]Slot

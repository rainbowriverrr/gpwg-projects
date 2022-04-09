package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Universe [][]bool

const (
	width  = 80
	height = 15
)

func mod(a, b int) int {
	return (a%b + b) % b
}

func NewUniverse() Universe {

	newUniverse := make(Universe, height)
	for i := range newUniverse {
		newUniverse[i] = make([]bool, width)
	}
	return newUniverse

}

func (u Universe) Show() {

	for i := 0; i < width; i++ {
		fmt.Printf("-")
	}
	fmt.Println()
	for y := range u {
		for x := range u[y] {
			if u[y][x] {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}

}

func (u Universe) Seed() {

	rand.Seed(time.Now().Unix())
	for y := range u {
		for x := range u[y] {
			num := rand.Intn(100)
			if num < 25 {
				u[y][x] = true
			}
		}
	}

}

func (u Universe) Alive(x, y int) bool {
	return u[mod(y, height)][mod(x, width)]
}

func (u Universe) Neighbors(x, y int) int {
	num := 0
	if u.Alive(x+1, y) {
		num++
	}
	if u.Alive(x-1, y) {
		num++
	}
	if u.Alive(x, y+1) {
		num++
	}
	if u.Alive(x, y-1) {
		num++
	}
	if u.Alive(x-1, y+1) {
		num++
	}
	if u.Alive(x-1, y-1) {
		num++
	}
	if u.Alive(x+1, y+1) {
		num++
	}
	if u.Alive(x+1, y-1) {
		num++
	}
	return num
}

func (u Universe) Next(x, y int) bool {

	numNeighbours := u.Neighbors(x, y)
	isAlive := false

	if u.Alive(x, y) {
		if numNeighbours == 2 || numNeighbours == 3 {
			isAlive = true
		}
	} else {
		if numNeighbours == 3 {
			isAlive = true
		}
	}

	return isAlive

}

func Step(curr, next Universe) {
	for y := range curr {
		for x := range curr[y] {
			next[y][x] = curr.Next(x, y)
		}
	}
}

func runSOL() {
	board := NewUniverse()
	board.Seed()
	var next Universe

	for i := 0; i < 50; i++ {
		fmt.Printf("\x1bc")
		next = NewUniverse()

		board.Show()

		Step(board, next)
		board = next
		time.Sleep(time.Millisecond * 250)
	}

}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type turtle struct {
	x int
	y int
}

func (t *turtle) move(x int, y int) {
	t.x += x
	t.y += y
}

func runTurtle() {

	rand.Seed(time.Now().Unix())

	player := &turtle{x: 0, y: 0}
	for i := 0; i < 10; i++ {

		fmt.Printf("Current position: %d, %d \n", player.x, player.y)

		x := rand.Intn(10)
		y := rand.Intn(10)

		fmt.Printf("Moving: %d, %d\n", x, y)

		player.move(x, y)

		time.Sleep(time.Second * 1)
	}

	fmt.Printf("Final position: %d, %d \n", player.x, player.y)

}

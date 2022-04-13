package main

import (
	"fmt"
	"strings"
)

//type MyString struct {
//	name string
//	last bool
//}

func runGoTest() {
	testString := "I really really love GO so so much"
	pipe1 := make(chan string)
	pipe2 := make(chan string)
	resultPipe := make(chan string)

	pipe1 <- testString
	go splitWords(pipe2, pipe1)
	close(pipe1)
	go delIdentical(resultPipe, pipe2)

	for word := range resultPipe {
		fmt.Print(word + " ")
	}

}

func splitWords(down chan string, up chan string) {
	for toSplit := range up { //Reads values from up until the channel is closed
		split := strings.Fields(toSplit)
		for _, curr := range split {
			down <- curr //sends words down
		}
	}
	//closes down once up is closed
	down <- ""
	close(down)
}

func delIdentical(down chan string, up chan string) {
	var prev string
	for curr := range up { //Reads from up until closed
		if curr == "" {
			close(down) //Closes down once final element of up is reached
			return
		} else {
			if curr != prev { //Sends string down if it is not identical to previous string
				down <- curr
			}
		}
		prev = curr
	}
}

package main

import (
	"strings"
)

//type MyString struct {
//	name string
//	last bool
//}

func runGoTest() {

}

func splitWords(down chan string, up chan string) {
	for toSplit := range up { //Reads values from up until the channel is closed
		split := strings.Fields(toSplit)
		for _, curr := range split {
			down <- curr //sends words down
		}
	}
	//closes down once up is closed
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

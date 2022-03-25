package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rainbowriverrr/gpwg-projects/sudoku"
)

func main() {

	fmt.Printf("\n River's 'Get Programming with Go' projects!")
	fmt.Printf("\n -------------------------------------------")
	fmt.Printf("\n These projects are capstone projects from ")
	fmt.Printf("\n the book.  Please select one of the options")
	fmt.Printf("\n below by entering it's number: \n")

	choosing := true
	for choosing {
		fmt.Printf("\n 1) Temperature Conversion table")
		fmt.Printf("\n 2) Simple Encoding/Decoding")
		fmt.Printf("\n 3) Game of Life Simulation")
		fmt.Printf("\n 4) Martian Animal Sancturary")
		fmt.Printf("\n 5) Turtle Moving")
		fmt.Printf("\n x) EXIT \n")

		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()

		if err != nil {
			fmt.Println(err)
		}

		switch char {
		case '1':
			fmt.Printf("\n\n Running Temperature Conversion Table...\n")
			runTemperature()
		case '2':
			fmt.Printf("\n\n Running Encoding/Decoding Example...\n")
			runEncoding()
		case '3':
			fmt.Printf("\n\n Running Game of Life Simulation...\n")
			runSOL()
		case '4':
			fmt.Printf("\n\n Running Martian Animal Sancturary...\n")
			runSancturary()
		case '5':
			fmt.Printf("\n\n Running Turtle Moving...\n")
			runTurtle()
		case '6':
			sudoku.TestSudoku()
		case 'x':
			os.Exit(0)
		default:
			fmt.Printf("\n\n Incorrect input, please try again \n\n")
		}
	}
}

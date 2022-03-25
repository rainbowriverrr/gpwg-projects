package main

import "fmt"

type Sudoku struct {
	board      [][]int8
	assigned   [][]bool
	candidates [][][]int8
}

func (s Sudoku) String() string {
	toReturn := ""
	for _, i := range s.board {
		for _, j := range i {
			toReturn += fmt.Sprint(j) + ","
		}
		toReturn += "\n"
	}
	return toReturn
}

func makeTestBoard() Sudoku {

	toReturn := Sudoku{}

	toReturn.board = [][]int8{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	toReturn.assigned = func() [][]bool {
		assigned := make([][]bool, 9)
		for i := range assigned {
			assigned[i] = make([]bool, 9)
			for j := range assigned[i] {
				if toReturn.board[i][j] == 0 {
					assigned[i][j] = false
				} else {
					assigned[i][j] = true
				}
			}
		}
		return assigned
	}()

	candidates := make([][][]int8, 9)
	for i := range candidates {
		candidates[i] = make([][]int8, 9)
	}

	toReturn.candidates = candidates

	return toReturn

}

//Only use when trying to generate candidates for a new Sudoku
func (s Sudoku) getAllCandidates() [][][]int8 {

	//Sets default list of candidates (none removed)
	candidates := make([][][]int8, 9)
	for i := range candidates {
		candidates[i] = make([][]int8, 9)
		for j := range candidates[i] {
			candidates[i][j] = []int8{1, 2, 3, 4, 5, 6, 7, 8, 9}
		}
	}

	for rowMultiplier := 0; rowMultiplier < 3; rowMultiplier++ { //Iterates through each 3x3 grid row-wise
		for colMultiplier := 0; colMultiplier < 3; colMultiplier++ { //Iterates through each 3x3 grid col-wise
			for row := 0 + (rowMultiplier)*3; row < 3+(rowMultiplier)*3; row++ { //Iterates through each row of the current grid
				for col := 0 + (colMultiplier)*3; col < 3+(colMultiplier)*3; col++ { //Iterates through each column of the current grid
					if s.board[row][col] != 0 {
						//Updates the candidate list of all spaces affected
						for i := 0; i < 9; i++ {
							for i, currCan := range candidates[row][i] {
								if currCan == s.board[row][col] {
									removeCandidate(&candidates[row][i], i)
								}
							}
						}
						for i := 0; i < 9; i++ {
							for i, currCan := range candidates[i][col] {
								if currCan == s.board[row][col] {
									removeCandidate(&candidates[i][col], i)
								}
							}
						}
					}
				}
			}
		}
	}

}

func (s Sudoku) set(x int, y int) error {

}

func (s Sudoku) clear(x int, y int) error {

}

func (s Sudoku) validateChoice(intent int, x int, y int) {

}

func removeCandidate(list *[]int8, index int) {
	if index < len(*list) { //moves last element into the place of the element to remove if the elment to remove is not the last element
		(*list)[index] = (*list)[len(*list)-1]
	}
	//removes last element in the list since that is now redundant
	*list = (*list)[:len(*list)-1]
}

/*
	This function updates all the candidates affected by the given space when a new number is inserted
	Input:
		(x int, y int): position of space which has been updated
*/
func (s *Sudoku) updateInsertedCandidate(x int, y int) {

	s.candidates[y][x] = []int8{}

	//Updates column
	for i := 0; i < 9; i++ {
		for i, currCan := range s.candidates[y][i] {
			if currCan == s.board[y][x] {
				removeCandidate(&s.candidates[y][i], i)
			}
		}
	}
	//Updates row
	for i := 0; i < 9; i++ {
		for i, currCan := range s.candidates[i][x] {
			if currCan == s.board[y][x] {
				removeCandidate(&s.candidates[i][x], i)
			}
		}
	}

	var rowOffset int
	var colOffset int

	if y < 3 {
		rowOffset = 0
	} else if y < 6 {
		rowOffset = 1
	} else {
		rowOffset = 2
	}
	if x < 3 {
		colOffset = 0
	} else if x < 6 {
		colOffset = 1
	} else {
		colOffset = 2
	}

	for row := 0 + rowOffset*3; row < 3+rowOffset*3; row++ {
		for col := 0 + colOffset*3; col < 3+colOffset*3; col++ {

		}
	}

}

func testSudoku() {
	testBoard := makeTestBoard()
	fmt.Print(testBoard)
}

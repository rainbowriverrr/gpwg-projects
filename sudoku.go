package main

import (
	"errors"
	"fmt"
)

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

	toReturn.candidates = func() [][][]int8 {
		candidates := make([][][]int8, 9)
		for i := range candidates {
			candidates[i] = make([][]int8, 9)
			for j := range candidates[i] {
				candidates[i][j] = []int8{1, 2, 3, 4, 5, 6, 7, 8, 9}
			}
		}
		return candidates
	}()

	return toReturn

}

/*
	input:
		x: column
		y: row
		toSet: number to set at position (x,y)
*/
func (s *Sudoku) set(x int8, y int8, toSet int8) error {
	if x < 0 || x >= 9 || y < 0 || y >= 9 {
		return errors.New("out of bounds")
	}
	if toSet < 1 || toSet > 9 {
		return errors.New("invalid number")
	}
	if s.assigned[y][x] {
		return errors.New("assigned number")
	}

	s.board[y][x] = toSet

	return nil
}

/*
	resets the number at the cartesian coordinate (x,y), returns error if out of bounds
	input:
		x: column
		y: row
*/
func (s *Sudoku) clear(x int8, y int8) error {
	if x < 0 || x >= 9 || y < 0 || y >= 9 {
		return errors.New("out of bounds")
	}
	if s.assigned[y][x] {
		return errors.New("assigned number")
	}

	s.board[y][x] = 0

	return nil
}

func testSudoku() {
	testBoard := makeTestBoard()
	fmt.Print(testBoard)
}

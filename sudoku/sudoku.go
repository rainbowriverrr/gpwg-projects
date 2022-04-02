package sudoku

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
	for currRow, i := range s.board {
		for currCol, j := range i {
			toReturn += fmt.Sprint(j) + " "
			if currCol == 2 || currCol == 5 {
				toReturn += "|"
			}
		}
		if currRow == 2 || currRow == 5 {
			toReturn += "\n--------------------"
		}
		toReturn += "\n"
	}
	return toReturn
}

/*
	if x exists within l, it is removed and the slice is shortened.
*/
func remove(l []int8, x int8) []int8 {
	if len(l) != 0 {

		var index int = -1

		for i := 0; i < len(l); i++ {
			if l[i] == x {
				index = i
			}
		}

		if index != -1 { //if element is found
			if index != len(l)-1 {
				l[index] = l[len(l)-1]
			}
			return l[:len(l)-1]
		}
	}
	//returns the original if l is empty or if the element is not found
	return l
}

func (s *Sudoku) updateAffectedCandidates(x int8, y int8) {

	s.candidates[y][x] = []int8{}

	for i := 0; i < 9; i++ {
		s.candidates[y][i] = remove(s.candidates[y][i], s.board[y][x])
	}
	for i := 0; i < 9; i++ {
		s.candidates[i][x] = remove(s.candidates[i][x], s.board[y][x])
	}

	var rowMin int
	var rowMax int
	var colMin int
	var colMax int

	if y < 3 {
		rowMin = 0
		rowMax = 3
	} else if y >= 3 && y < 6 {
		rowMin = 3
		rowMax = 6
	} else {
		rowMin = 6
		rowMax = 9
	}

	if x < 3 {
		colMin = 0
		colMax = 3
	} else if x >= 3 && x < 6 {
		colMin = 3
		colMax = 6
	} else {
		colMin = 6
		colMax = 9
	}

	for currRow := rowMin; currRow < rowMax; currRow++ {
		for currCol := colMin; currCol < colMax; currCol++ {
			if currCol != int(x) && currRow != int(y) {
				s.candidates[currRow][currCol] = remove(s.candidates[currRow][currCol], s.board[y][x])
			}
		}
	}
}

func (s *Sudoku) setAllCandidates() {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if s.board[row][col] != 0 {
				s.updateAffectedCandidates(int8(col), int8(row))
			}
		}
	}
}

func (s *Sudoku) Solve() {

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

	toReturn.setAllCandidates()

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

	valid := false
	for _, candidate := range s.candidates[y][x] {
		if candidate == toSet {
			valid = true
		}
	}

	if !valid {
		return errors.New("breaks rules")
	}

	s.board[y][x] = toSet
	s.updateAffectedCandidates(x, y)

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
	s.updateAffectedCandidates(x, y)

	return nil
}

func TestSudoku() {
	testBoard := makeTestBoard()
	fmt.Println(testBoard)

	err := testBoard.set(2, 0, 4)

	if err != nil {
		fmt.Println(err)
	}

	err = testBoard.set(3, 0, 4)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(testBoard)

	testBoard.Solve()

	fmt.Println(testBoard.candidates)
}

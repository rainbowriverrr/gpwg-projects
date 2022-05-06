package sudoku

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
	//"math/rand"
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

//TODO:  Make algorithm reset candidates for clear
func (s *Sudoku) updateAffectedCandidates(x int8, y int8) {

	//Clear
	if s.board[y][x] == 0 {

		//resets candidates for empty spaces first
		for i := 0; i < 9; i++ {
			if s.board[y][i] == 0 {
				s.candidates[y][i] = []int8{1, 2, 3, 4, 5, 6, 7, 8, 9}
			}
			if s.board[x][i] == 0 {
				s.candidates[i][x] = []int8{1, 2, 3, 4, 5, 6, 7, 8, 9}
			}
		}

		s.setAllCandidates()

	} else { //Set

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

/*
	returns(row, column)
	Finds the last empty cell,  checks if the sudoku is solvable,
	-1,-1 is returned if unsolvable
	10,10 is returned if no empties (solved)
*/
func (s *Sudoku) findBestEmpty() (int8, int8) {

	minCandidates := 9
	var minCol int8 = 10
	var minRow int8 = 10

	var row int8
	var col int8

	for row = 0; row < 9; row++ {
		for col = 0; col < 9; col++ {
			if s.board[row][col] == 0 {
				if len(s.candidates[row][col]) == 0 {
					return -1, -1
				} else {
					if len(s.candidates[row][col]) <= minCandidates {
						minCandidates = len(s.candidates[row][col])
						minCol = col
						minRow = row
					}
				}
			}
		}
	}

	return minRow, minCol

}

func (s *Sudoku) Solve() bool {

	time.Sleep(time.Millisecond * 500)
	emptyRow, emptyCol := s.findBestEmpty()
	fmt.Println(emptyRow, emptyCol)
	fmt.Println(s)
	if emptyRow == 10 {
		return true
	} else if emptyCol == -1 {
		return false
	}

	nextBoard := s

	for _, num := range s.candidates[emptyRow][emptyCol] {
		nextBoard.set(emptyCol, emptyRow, num)
		if nextBoard.Solve() {
			return true
		} else {
			nextBoard.clear(emptyCol, emptyRow)
		}
	}

	return false

}

func (s *Sudoku) SolveRandom() bool {
	time.Sleep(time.Millisecond * 500)
	emptyRow, emptyCol := s.findBestEmpty()
	fmt.Println(emptyRow, emptyCol)
	fmt.Println(s)

	if emptyRow == 10 {
		return true
	} else if emptyCol == -1 {
		return false
	}

	for initialSize := len(s.candidates[emptyRow][emptyCol]); initialSize > 0; initialSize++ {
		number := s.candidates[emptyRow][emptyCol][rand.Intn(len(s.candidates[emptyRow][emptyCol]))]
		s.set(emptyCol, emptyRow, number)
		if s.Solve() {
			return true
		} else {
			s.clear(emptyCol, emptyRow)
		}
	}

	return false

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

func GenerateSudoku(difficulty int) Sudoku {
	toReturn := Sudoku{}

	//Generate empty grid
	toReturn.board = func() [][]int8 {
		assigned := make([][]int8, 9)
		for i := range assigned {
			assigned[i] = make([]int8, 9)
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

	toReturn.assigned = func() [][]bool {
		assigned := make([][]bool, 9)
		for i := range assigned {
			assigned[i] = make([]bool, 9)
			for j := range assigned[i] {
				assigned[i][j] = false
			}
		}
		return assigned
	}()

	//Fill empty grid by solving it
	rand.Seed(time.Now().UnixNano())
	toReturn.SolveRandom()

	//Remove values based on difficulty
	toRemove := rand.Intn(5)
	if difficulty <= 0 {
		//easy 22 - 26
		toRemove += 22
	} else if difficulty == 1 {
		//medium 36 - 40
		toRemove += 36
	} else if difficulty == 2 {
		//hard 44 - 48
		toRemove += 44
	} else if difficulty == 3 {
		//difficult 50 - 54
		toRemove += 50
	} else {
		//really difficulty 54 - 58
		toRemove += 54
	}
	var col int = rand.Intn(9)
	var row int = rand.Intn(9)
	for toRemove > 0 {

		for toReturn.board[row][col] == 0 {
			col = rand.Intn(9)
			row = rand.Intn(9)
		}

		toReturn.clear(int8(col), int8(row))

		toRemove--
	}
	//Set 'Assigned' field
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

	return toReturn
}

func TestSudoku() {
	testBoard := GenerateSudoku(2)

	fmt.Println(testBoard)
	//
	//	fmt.Println(testBoard.candidates)
	//
	//	err := testBoard.set(2, 0, 4)
	//
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//
	//	fmt.Println(testBoard)
	//
	//	fmt.Println(testBoard.candidates)
	//
	//	err = testBoard.clear(2, 0)
	//
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//
	//	fmt.Println(testBoard)
	//
	//	fmt.Println(testBoard.candidates)

	//fmt.Println(testBoard.SolveRandom())
}

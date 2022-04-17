package lifeonmars

import (
	"image"
)

//represents an area on the surface of Mars
//may be used concurrently by goroutines
type MarsGrid struct {
	//TODO
	maxX      int
	maxY      int
	minX      int
	minY      int
	Occupiers []Occupier
}

//Represents an occupied cell on the grid.
//May be used concurrently by different goroutines
type Occupier struct {
	//TODO
	point image.Point
}

type RoverDriver struct {
	Occupier
}

//Occupies a cell at the given point in the grid.
//Returns nil if the point is already occupied, or the point is outside of the grid
//Otherwise, returns a value that can be used to move to different places on the grid
func (g *MarsGrid) Occupy(p image.Point) *Occupier

//Moves the occupier to a different cell in the grid.
//Returns whether the move was successful
//Could fail from invalid move order, or because point is already occupied
//Occupier remains in the same place if this fails
func (o *Occupier) Move(p image.Point) bool

func testLOM() {

}

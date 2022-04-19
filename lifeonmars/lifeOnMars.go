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
	Occupiers []*Occupier
}

//Represents an occupied cell on the grid.
//May be used concurrently by different goroutines
type Occupier struct {
	//TODO
	point image.Point
	grid  *MarsGrid
}

type RoverDriver struct {
	Occupier
}

func newRoverDriver(g *MarsGrid) *RoverDriver {
	toReturn := &RoverDriver{}
	toReturn.grid = g

	return toReturn
}

//Returns pointer to MarsGrid of size 100x100 with no Occupiers
func newMarsGrid() *MarsGrid {
	toReturn := &MarsGrid{
		maxX:      50,
		maxY:      50,
		minX:      -50,
		minY:      -50,
		Occupiers: make([]*Occupier, 10),
	}

	return toReturn
}

//Occupies a cell at the given point in the grid.
//Returns nil if the point is already occupied, or the point is outside of the grid
//Otherwise, returns a value that can be used to move to different places on the grid
func (g *MarsGrid) Occupy(p image.Point) *Occupier {
	if p.X < g.minX || p.X > g.maxX {
		return nil
	}
	if p.Y < g.minY || p.Y > g.maxY {
		return nil
	}
	for _, curr := range g.Occupiers {
		if curr.point.X == p.X && curr.point.Y == p.Y {
			return nil
		}
	}
	return &Occupier{
		point: p,
		grid:  g,
	}
}

//Moves the occupier to a different cell in the grid.
//Returns whether the move was successful
//Could fail from invalid move order, or because point is already occupied
//Occupier remains in the same place if this fails
func (o *Occupier) Move(p image.Point) bool {
	if p.X < o.grid.minX || p.X > o.grid.maxX {
		return false
	}
	if p.Y < o.grid.minY || p.Y > o.grid.maxY {
		return false
	}
	for _, curr := range o.grid.Occupiers {
		if curr.point.X == p.X && curr.point.Y == p.Y {
			return false
		}
	}
	o.point = p
	return true
}

func testLOM() {

}

package main

import (
	"fmt"
	"io"
	"math"
	"strings"
)

const (
	Top    = 0
	Left   = 1
	Right  = 2
	Bottom = 3

	TopLeft     = 4
	TopRight    = 5
	BottomLeft  = 6
	BottomRight = 7
)

var (
	Opposite = [...]Cell{
		Top:    Bottom,
		Left:   Right,
		Right:  Left,
		Bottom: Top,
	}
)

// Coord Cartesian coordinate
type Coord struct {
	X, Y int
}

func (c Coord) Equal(cmp Coord) bool {
	return c.X == cmp.X && c.Y == cmp.Y
}

func (c Coord) String() string {
	var sb strings.Builder
	sb.Grow(5)
	sb.WriteString(intToStr(c.X))
	sb.WriteString(".")
	sb.WriteString(intToStr(c.Y))
	return sb.String()
}

// Cell index system for a 2D Grid numbered Left-Right Top-Bottom
type (
	Cell  = int
	Cells []Cell
)

// Grid Implementation of a 2D field using indexes instead of Cartesian
type Grid struct {
	Width  int
	Height int
	Size   int
	Cells  Cells

	ByCell  []Coord
	ByCoord []Cells

	Moves []Cells

	Center Cell
}

func NewGrid(width, height int) *Grid {
	size := height * width

	grid := &Grid{
		Width:  width,
		Height: height,
		Size:   size,

		Cells: make(Cells, 0, size),

		ByCell:  make([]Coord, 0, size),
		ByCoord: make([]Cells, width),

		Moves: make([]Cells, size),
	}

	for id := 0; id < size; id++ {
		grid.Cells = append(grid.Cells, id)
		for _, aid := range grid.CellSides(id) {
			if aid == -1 {
				continue
			}
			grid.Moves[id] = append(grid.Moves[id], aid)
		}
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			grid.ByCell = append(grid.ByCell, Coord{x, y})
		}
	}
	for cell, coord := range grid.ByCell {
		grid.ByCoord[coord.X] = append(grid.ByCoord[coord.X], cell)
	}

	grid.Center = size / 2

	return grid
}

func NewGridFromInput(r io.Reader) *Grid {
	var width, height int
	n, err := fmt.Fscanln(r, &width, &height)
	if err != nil {
		panic(err)
	}
	if n != 2 {
		panic(fmt.Errorf(`invalid input arguments %d for size expected 2`, n))
	}

	return NewGrid(width, height)
}

func (g *Grid) CellExists(id Cell) bool {
	if id < 0 || id >= g.Size {
		return false
	}
	return true
}

func (g *Grid) CoordExists(id Coord) bool {
	if id.Y < 0 || id.Y >= g.Height {
		return false
	}
	if id.X < 0 || id.X >= g.Width {
		return false
	}
	return true
}

func (g *Grid) ToCoord(id Cell) Coord {
	return g.ByCell[id]
}

func (g *Grid) ToCell(id Coord) Cell {
	return g.ByCoord[id.X][id.Y]
}

func (g *Grid) DistanceManhattan(from, to Cell) int {
	return abs(g.ByCell[from].X-g.ByCell[to].X) + abs(g.ByCell[from].Y-g.ByCell[to].Y)
}

func (g *Grid) DistanceEuclidean(from, to Cell) float64 {
	return math.Sqrt(math.Pow(float64(g.ByCell[from].X-g.ByCell[to].X), 2) + math.Pow(float64(g.ByCell[from].Y-g.ByCell[to].Y), 2))
}

func (g *Grid) Top(id Cell) Cell {
	if id-g.Width >= 0 {
		return id - g.Width
	}
	return -1
}

func (g *Grid) TopLeft(id Cell) Cell {
	top := g.Top(id)
	if top == -1 {
		return -1
	}
	return g.Left(top)
}

func (g *Grid) TopRight(id Cell) Cell {
	top := g.Top(id)
	if top == -1 {
		return -1
	}
	return g.Right(top)
}

func (g *Grid) Left(id Cell) Cell {
	if id%g.Width != 0 {
		return id - 1
	}
	return -1
}

func (g *Grid) Right(id Cell) Cell {
	if id%g.Width != g.Width-1 {
		return id + 1
	}
	return -1
}

func (g *Grid) Bottom(id Cell) Cell {
	if id+g.Width < g.Size {
		return id + g.Width
	}
	return -1
}

func (g *Grid) BottomLeft(id Cell) Cell {
	bottom := g.Bottom(id)
	if bottom == -1 {
		return -1
	}
	return g.Left(bottom)
}

func (g *Grid) BottomRight(id Cell) Cell {
	bottom := g.Bottom(id)
	if bottom == -1 {
		return -1
	}
	return g.Right(bottom)
}

func (g *Grid) CellSides(id Cell) Cells {
	return Cells{
		Top:    g.Top(id),
		Left:   g.Left(id),
		Right:  g.Right(id),
		Bottom: g.Bottom(id),
	}
}

func (g *Grid) CellAdjacent(id Cell) Cells {
	return Cells{
		Top:         g.Top(id),
		Left:        g.Left(id),
		Right:       g.Right(id),
		Bottom:      g.Bottom(id),
		TopLeft:     g.TopLeft(id),
		TopRight:    g.TopRight(id),
		BottomLeft:  g.BottomLeft(id),
		BottomRight: g.BottomRight(id),
	}
}

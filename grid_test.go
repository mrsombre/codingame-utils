package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testIdFirst = 0
	testId      = 12
	testIdLast  = 24

	testGrid = func(params ...int) *Grid {
		cnt := len(params)
		if cnt != 0 && cnt != 2 {
			panic(fmt.Errorf(`invalid count of arguments %d`, cnt))
		}
		var width, height int
		if cnt == 0 {
			width = 5
			height = 5
		} else {
			width = params[0]
			height = params[1]
		}
		return NewGrid(width, height)
	}
)

func TestCoord_Equal(t *testing.T) {
	tests := []struct {
		name     string
		a        Coord
		b        Coord
		expected bool
	}{
		{
			name:     `equal`,
			a:        Coord{1, 1},
			b:        Coord{1, 1},
			expected: true,
		},
		{
			name:     `not equal`,
			a:        Coord{1, 1},
			b:        Coord{0, 1},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.a.Equal(tc.b))
		})
	}
}

func TestCoord_String(t *testing.T) {
	tests := []struct {
		name     string
		coord    Coord
		expected string
	}{
		{
			name:     `1.2`,
			coord:    Coord{1, 2},
			expected: `1.2`,
		},
		{
			name:     `12.34`,
			coord:    Coord{12, 34},
			expected: `12.34`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.coord.String())
		})
	}
}

func TestNewGrid(t *testing.T) {
	grid := NewGrid(3, 3)

	assert.Equal(t, 9, grid.Size)
	cells := Cells{0, 1, 2, 3, 4, 5, 6, 7, 8}
	assert.Equal(t, cells, grid.Cells)

	byCell := []Coord{{0, 0}, {1, 0}, {2, 0}, {0, 1}, {1, 1}, {2, 1}, {0, 2}, {1, 2}, {2, 2}}
	assert.Equal(t, byCell, grid.ByCell)

	byCoord := []Cells{{0, 3, 6}, {1, 4, 7}, {2, 5, 8}}
	assert.Equal(t, byCoord, grid.ByCoord)

	assert.Len(t, grid.Moves, 9)
	assert.Equal(t, Cells{1, 3}, grid.Moves[0])

	assert.Equal(t, 4, grid.Center)
}

func TestNewGridFromInput(t *testing.T) {
	input := "4 3\n"
	grid := NewGridFromInput(bytes.NewBufferString(input))

	assert.Equal(t, 12, grid.Size)
}

func TestGrid_CellExists(t *testing.T) {
	tests := []struct {
		name     string
		cell     Cell
		expected bool
	}{
		{
			name:     `first`,
			cell:     testIdFirst,
			expected: true,
		},
		{
			name:     `center`,
			cell:     testId,
			expected: true,
		},
		{
			name:     `last`,
			cell:     testIdLast,
			expected: true,
		},
		{
			name:     `-1`,
			cell:     testIdFirst - 1,
			expected: false,
		},
		{
			name:     `last+1`,
			cell:     testIdLast + 1,
			expected: false,
		},
	}

	grid := testGrid()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, grid.CellExists(tc.cell))
		})
	}
}

func TestGrid_CoordExists(t *testing.T) {
	tests := []struct {
		name     string
		coord    Coord
		expected bool
	}{
		{
			name:     `first`,
			coord:    Coord{0, 0},
			expected: true,
		},
		{
			name:     `center`,
			coord:    Coord{2, 2},
			expected: true,
		},
		{
			name:     `last`,
			coord:    Coord{4, 4},
			expected: true,
		},
		{
			name:     `-1`,
			coord:    Coord{-1, -1},
			expected: false,
		},
		{
			name:     `+1`,
			coord:    Coord{6, 6},
			expected: false,
		},
	}

	grid := testGrid()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, grid.CoordExists(tc.coord))
		})
	}
}

func TestGrid_ToCell(t *testing.T) {
	assert.Equal(t, testId, testGrid().ToCell(Coord{2, 2}))
}

func TestGrid_ToCoord(t *testing.T) {
	assert.Equal(t, Coord{2, 2}, testGrid().ToCoord(testId))
}

func TestGrid_DistanceManhattan(t *testing.T) {
	tests := []struct {
		name     string
		from, to Cell
		expected int
	}{
		{
			name:     `1`,
			from:     12,
			to:       13,
			expected: 1,
		},
		{
			name:     `1`,
			from:     0,
			to:       12,
			expected: 4,
		},
	}

	grid := testGrid()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, grid.DistanceManhattan(tc.from, tc.to))
		})
	}
}

func TestGrid_DistanceEuclidean(t *testing.T) {
	tests := []struct {
		name     string
		from, to Cell
		expected float64
	}{
		{
			name:     `10`,
			from:     10,
			to:       12,
			expected: 2,
		},
		{
			name:     `16`,
			from:     16,
			to:       12,
			expected: 1.4,
		},
		{
			name:     `22`,
			from:     22,
			to:       12,
			expected: 2,
		},
	}

	grid := testGrid()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.InDelta(t, tc.expected, grid.DistanceEuclidean(tc.from, tc.to), 0.5)
		})
	}
}

func TestGrid_CellSides(t *testing.T) {
	tests := []struct {
		name     string
		cell     Cell
		expected Cells
	}{
		{
			name:     `first`,
			cell:     0,
			expected: Cells{-1, -1, 1, 4},
		},
		{
			name:     `center`,
			cell:     5,
			expected: Cells{1, 4, 6, 9},
		},
		{
			name:     `last`,
			cell:     11,
			expected: Cells{7, 10, -1, -1},
		},
		{
			name:     `corner left-bottom`,
			cell:     8,
			expected: Cells{4, -1, 9, -1},
		},
		{
			name:     `corner right-top`,
			cell:     3,
			expected: Cells{-1, 2, -1, 7},
		},
		{
			name:     `left`,
			cell:     4,
			expected: Cells{0, -1, 5, 8},
		},
		{
			name:     `right`,
			cell:     7,
			expected: Cells{3, 6, -1, 11},
		},
		{
			name:     `top`,
			cell:     2,
			expected: Cells{-1, 1, 3, 6},
		},
		{
			name:     `bottom`,
			cell:     9,
			expected: Cells{5, 8, 10, -1},
		},
	}

	/*
		0  1  2  3
		4  5  6  7
		8  9  10 11
	*/
	grid := testGrid(4, 3)
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, grid.CellSides(tc.cell))
		})
	}
}

func TestGrid_CellAdjacent(t *testing.T) {
	tests := []struct {
		name     string
		cell     Cell
		expected Cells
	}{
		{
			name:     `center`,
			cell:     5,
			expected: Cells{1, 4, 6, 9, 0, 2, 8, 10},
		},
		{
			name:     `corner left-bottom`,
			cell:     8,
			expected: Cells{4, -1, 9, -1, -1, 5, -1, -1},
		},
		{
			name:     `corner right-top`,
			cell:     3,
			expected: Cells{-1, 2, -1, 7, -1, -1, 6, -1},
		},
	}

	/*
		0  1  2  3
		4  5  6  7
		8  9  10 11
	*/
	grid := testGrid(4, 3)
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, grid.CellAdjacent(tc.cell))
		})
	}
}

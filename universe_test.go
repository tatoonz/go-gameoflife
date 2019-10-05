package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextLifeOfCell_ShouldDie_WhenNeighbourLivesLessThanTwo(t *testing.T) {
	u := &universe{
		rows:    3,
		columns: 3,
		cells: [][]bool{
			{false, false, false},
			{false, true, false},
			{false, false, false},
		},
	}

	assert.False(t, u.nextLifeOfCell(1, 1))
}

func TestNextLifeOfCell_ShouldDie_WhenAliveCellHasMoreThan3LiveNeighbours(t *testing.T) {
	u := &universe{
		rows:    3,
		columns: 3,
		cells: [][]bool{
			{true, false, false},
			{false, true, true},
			{false, true, true},
		},
	}

	assert.False(t, u.nextLifeOfCell(1, 1))
}

func TestNextLifeOfCell_ShouldStillAlive_WhenCellIsAlive(t *testing.T) {
	t.Run("AndHas2LiveNeighbour", func(t *testing.T) {
		u := &universe{
			rows:    3,
			columns: 3,
			cells: [][]bool{
				{true, false, false},
				{false, true, true},
				{false, false, false},
			},
		}

		assert.True(t, u.nextLifeOfCell(1, 1))
	})

	t.Run("AndHas3LiveNeighbour", func(t *testing.T) {
		u := &universe{
			rows:    3,
			columns: 3,
			cells: [][]bool{
				{true, false, true},
				{false, true, true},
				{false, false, false},
			},
		}

		assert.True(t, u.nextLifeOfCell(1, 1))
	})
}

func TestNextLifeOfCell_ShouldResurrect_WhenDeadCellHasExact3LiveNeighbours(t *testing.T) {
	u := &universe{
		rows:    3,
		columns: 3,
		cells: [][]bool{
			{true, false, false},
			{false, false, true},
			{false, true, false},
		},
	}

	assert.True(t, u.nextLifeOfCell(1, 1))
}

func TestNextLifeOfCell_ShouldStillDead_WhenDeadCellHasNotExact3LiveNeighbours(t *testing.T) {
	t.Run("LessThanCase", func(t *testing.T) {
		u := &universe{
			rows:    3,
			columns: 3,
			cells: [][]bool{
				{true, false, false},
				{false, false, false},
				{false, true, false},
			},
		}

		assert.False(t, u.nextLifeOfCell(1, 1))
	})

	t.Run("MoreThanCase", func(t *testing.T) {
		u := &universe{
			rows:    3,
			columns: 3,
			cells: [][]bool{
				{true, false, true},
				{false, false, true},
				{false, true, false},
			},
		}

		assert.False(t, u.nextLifeOfCell(1, 1))
	})
}

func TestLiveNeighboursOfCell_ShouldReturnCorrectNumber_ForNormalCase(t *testing.T) {
	u := &universe{
		rows:    3,
		columns: 3,
		cells: [][]bool{
			{true, true, true},
			{true, true, true},
			{true, true, true},
		},
	}

	assert.Equal(t, 8, u.liveNeighboursOfCell(1, 1))
}

func TestLiveNeighboursOfCell_ShouldReturnCorrectNumber_ForLeftEdgeCase(t *testing.T) {
	u := &universe{
		rows:    3,
		columns: 3,
		cells: [][]bool{
			{false, false, true},
			{true, false, true},
			{false, false, true},
		},
	}

	assert.Equal(t, 3, u.liveNeighboursOfCell(0, 1))
}

func TestLiveNeighboursOfCell_ShouldReturnCorrectNumber_ForRightEdgeCase(t *testing.T) {
	u := &universe{
		rows:    3,
		columns: 3,
		cells: [][]bool{
			{true, false, false},
			{true, false, true},
			{true, false, false},
		},
	}

	assert.Equal(t, 3, u.liveNeighboursOfCell(2, 1))
}

func TestLiveNeighboursOfCell_ShouldReturnCorrectNumber_ForTopEdgeCase(t *testing.T) {
	u := &universe{
		rows:    3,
		columns: 3,
		cells: [][]bool{
			{false, true, false},
			{false, false, false},
			{true, true, true},
		},
	}

	assert.Equal(t, 3, u.liveNeighboursOfCell(1, 0))
}

func TestLiveNeighboursOfCell_ShouldReturnCorrectNumber_ForBottomEdgeCase(t *testing.T) {
	u := &universe{
		rows:    3,
		columns: 3,
		cells: [][]bool{
			{true, true, true},
			{false, false, false},
			{false, true, false},
		},
	}

	assert.Equal(t, 3, u.liveNeighboursOfCell(1, 2))
}

func TestLiveNeighboursOfCell_ShouldReturnCorrectNumber_ForCornerCases(t *testing.T) {
	u := &universe{
		rows:    3,
		columns: 3,
		cells: [][]bool{
			{true, false, true},
			{false, false, false},
			{true, false, true},
		},
	}

	t.Run("TopLeft", func(t *testing.T) {
		assert.Equal(t, 3, u.liveNeighboursOfCell(0, 0))
	})

	t.Run("BottomLeft", func(t *testing.T) {
		assert.Equal(t, 3, u.liveNeighboursOfCell(0, 2))
	})

	t.Run("TopRight", func(t *testing.T) {
		assert.Equal(t, 3, u.liveNeighboursOfCell(2, 0))
	})

	t.Run("BottomRight", func(t *testing.T) {
		assert.Equal(t, 3, u.liveNeighboursOfCell(2, 2))
	})
}

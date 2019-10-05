package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextLifeOfCell_ShouldReturnFalse_WhenNeighbourLivesLessThanTwo(t *testing.T) {
	u := &universe{
		cells: [][]bool{
			{false, false, false},
			{false, true, false},
			{false, false, false},
		},
	}

	assert.False(t, u.nextLifeOfCell(1, 1))
}

func TestNextLifeOfCell_ShouldStillAlive_WhenCellIsAlive(t *testing.T) {
	t.Run("AndHas2LiveNeighbour", func(t *testing.T) {
		u := &universe{
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
			cells: [][]bool{
				{true, false, true},
				{false, true, true},
				{false, false, false},
			},
		}

		assert.True(t, u.nextLifeOfCell(1, 1))
	})
}

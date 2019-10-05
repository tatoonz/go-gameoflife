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

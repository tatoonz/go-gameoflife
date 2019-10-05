package main

import (
	"math/rand"
	"time"
)

type universe struct {
	rows  int
	cols  int
	cells [][]bool
}

func newUniverse(rows, cols int) *universe {
	cells := make([][]bool, rows)
	for y := range cells {
		cells[y] = make([]bool, cols)
	}

	return &universe{rows: rows, cols: cols, cells: cells}
}

func (u *universe) next() *universe {
	nextU := newUniverse(u.rows, u.cols)

	for y := range u.cells {
		for x := range u.cells[y] {
			nextU.cells[y][x] = u.nextLifeOfCell(x, y)
		}
	}

	return nextU
}

func (u *universe) cellAlive(x, y int) bool {
	return u.cells[y][x]
}

func (u *universe) nextLifeOfCell(x, y int) bool {
	liveNeighBours := u.liveNeighboursOfCell(x, y)
	cellAlive := u.cellAlive(x, y)

	if cellAlive && (liveNeighBours == 2 || liveNeighBours == 3) {
		return true
	}

	if !cellAlive && liveNeighBours == 3 {
		return true
	}

	return false
}

func (u *universe) liveNeighboursOfCell(x, y int) int {
	result := 0

	prevX := x - 1
	if prevX < 0 {
		prevX = u.cols - 1
	}

	nextX := x + 1
	if nextX > u.cols-1 {
		nextX = 0
	}

	prevY := y - 1
	if prevY < 0 {
		prevY = u.rows - 1
	}

	nextY := y + 1
	if nextY > u.rows-1 {
		nextY = 0
	}

	// top left
	if u.cellAlive(prevX, prevY) {
		result++
	}

	// left
	if u.cellAlive(prevX, y) {
		result++
	}

	// bottom left
	if u.cellAlive(prevX, nextY) {
		result++
	}

	// top
	if u.cellAlive(x, prevY) {
		result++
	}

	// bottom
	if u.cellAlive(x, nextY) {
		result++
	}

	// top right
	if u.cellAlive(nextX, prevY) {
		result++
	}

	// right
	if u.cellAlive(nextX, y) {
		result++
	}

	// bottom right
	if u.cellAlive(nextX, nextY) {
		result++
	}

	return result
}

func (u *universe) randCells() {
	rand.Seed(time.Now().UnixNano())

	for y := 0; y < u.rows-1; y++ {
		for x := 0; x < u.cols-1; x++ {
			u.cells[y][x] = rand.Intn(2) == 0
		}
	}
}

package main

import "fmt"

type universe struct {
	rows    int
	columns int
	cells   [][]bool
}

func (u *universe) cellAlive(x, y int) bool {
	return u.cells[y][x]
}

func (u *universe) nextLifeOfCell(x, y int) bool {
	liveNeighBours := u.countNeighbourLivesOfCell(x, y)

	fmt.Println("liveNeighBours", liveNeighBours)

	if liveNeighBours == 2 || liveNeighBours == 3 {
		return true
	}

	return false
}

func (u *universe) countNeighbourLivesOfCell(x, y int) int {
	result := 0

	prevX := x - 1
	if prevX < 0 {
		prevX = u.columns - 1
	}

	nextX := x + 1
	if nextX > u.columns-1 {
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

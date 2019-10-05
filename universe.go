package main

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

	// top left
	if u.cellAlive(prevX, y-1) {
		result++
	}

	// left
	if u.cellAlive(prevX, y) {
		result++
	}

	// bottom left
	if u.cellAlive(prevX, y+1) {
		result++
	}

	// top
	if u.cellAlive(x, y-1) {
		result++
	}

	// bottom
	if u.cellAlive(x, y+1) {
		result++
	}

	// top right
	if u.cellAlive(x+1, y-1) {
		result++
	}

	// right
	if u.cellAlive(x+1, y) {
		result++
	}

	// bottom right
	if u.cellAlive(x+1, y+1) {
		result++
	}

	return result
}

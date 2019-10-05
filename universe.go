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
	liveNeighBours := 0
	// top left
	if u.cellAlive(x-1, y-1) {
		liveNeighBours++
	}

	// left
	if u.cellAlive(x-1, y) {
		liveNeighBours++
	}

	// bottom left
	if u.cellAlive(x-1, y+1) {
		liveNeighBours++
	}

	// top
	if u.cellAlive(x, y-1) {
		liveNeighBours++
	}

	// bottom
	if u.cellAlive(x, y+1) {
		liveNeighBours++
	}

	// top right
	if u.cellAlive(x+1, y-1) {
		liveNeighBours++
	}

	// right
	if u.cellAlive(x+1, y) {
		liveNeighBours++
	}

	// bottom right
	if u.cellAlive(x+1, y+1) {
		liveNeighBours++
	}

	if liveNeighBours == 2 || liveNeighBours == 3 {
		return true
	}

	return false
}

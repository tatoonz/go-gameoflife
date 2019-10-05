package main

type universe struct {
	rows  int
	cols  int
	cells [][]bool
}

func newCells(rows, cols int) [][]bool {
	result := make([][]bool, cols)

	for y := range result {
		result[y] = make([]bool, rows)
	}

	return result
}

func (u *universe) next() *universe {
	nextU := &universe{rows: u.rows, cols: u.cols, cells: newCells(u.rows, u.cols)}

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

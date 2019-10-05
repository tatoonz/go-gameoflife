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
	return false
}

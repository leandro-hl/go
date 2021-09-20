package lifegame

// Field represents a two-dimensional field of cells.
type Field struct {
	s    [][]bool
	w, h int
}

// Life stores the state of a round of Conway's Game of Life.
type Life struct {
	a, b *Field
	w, h int
}

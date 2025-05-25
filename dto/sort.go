package dto

type Direction int

type Sort struct {
	Field string
	Value Direction
}

const (
	ASC  Direction = 1
	DESC Direction = -1
)

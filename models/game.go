package models

type Light struct {
	Level uint8
	Color uint8
}

type Position struct {
	X uint16
	Y uint16
	Z uint8
}

const (
	North uint8 = iota
	East
	South
	West
)

type Offset struct {
	X, Y, Z int8
}

func (pos *Position) Offset(offset Offset) {
	pos.X += (uint16)(offset.X)
	pos.Y += (uint16)(offset.Y)
	pos.Z += (uint8)(offset.Z)
}

package flag

type Flag int8

const (
	Unknown Flag = -1
	None         = iota
	Green
	Blue
	Yellow
	Red
)

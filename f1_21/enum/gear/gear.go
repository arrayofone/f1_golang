package gear

type Gear int8

// Gear selected (1-8, N=0, R=-1)
const (
	R Gear = -1
	N      = iota
)

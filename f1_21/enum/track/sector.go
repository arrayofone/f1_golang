package track

type Sector uint8

// 0 = sector1, 1 = sector2, 2 = sector3

const (
	Sector1 Sector = iota
	Sector2
	Sector3
)

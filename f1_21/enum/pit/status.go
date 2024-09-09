package pit

type PitStatus uint8

// 0 = none, 1 = pitting, 2 = in pit area

const (
	None PitStatus = iota
	Pitting
	InPitArea
)

package driver

type ResultStatus uint8

// 0 = invalid, 1 = inactive, 2 = active 3 = finished, 4 = didnotfinish, 5 = disqualified 6 = not classified, 7 = retired

const (
	Invalid ResultStatus = iota
	Inactive
	Active
	Finished
	DNF
	Disqualified
	NoClassified
	Retired
)

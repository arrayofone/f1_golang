package tyres

type Compound uint8

// F1 Modern -
// 16 = C5,
// 17 = C4,
// 18 = C3,
// 19 = C2,
// 20 = C1
// 7 = inter,
// 8 = wet
// F1 Classic -
// 9 = dry,
// 10 = wet
// F2 –
// 11 = super soft,
// 12 = soft,
// 13 = medium,
// 14 = hard
// 15 = wet

const (
	_ Compound = iota
	_
	_
	_
	_
	_
	_
	Inter
	Wet
	DryClassic
	WetClassic
	SuperSoftF2
	SoftF2
	MediumF2
	HardF2
	WetF2
	C5
	C4
	C3
	C2
	C1
)

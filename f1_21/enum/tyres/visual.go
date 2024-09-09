package tyres

type Visual uint8

// F1 visual (can be different from actual compound)
// 16 = soft,
// 17 = medium,
// 18 = hard,
// 7 = inter,
// 8 = wet
// F1 Classic -
// 9 = dry,
// 10 = wet,
// F2 ‘19,
// 15 = wet,
// 19 – super soft,
// 20 = soft,
// 21 = medium ,
// 22 = hard

const (
	_ Visual = iota
	_
	_
	_
	_
	_
	_
	InterVisual
	WetVisual
	DryVisualClassic
	WetVisualClassic
	_
	_
	_
	_
	WetVisualF2
	SoftVisualF1
	MediumVisualF1
	HardVisualF1
	SuperSoftVisualF2
	SoftVisualF2
	MediumVisualF2
	HardVisualF2
)

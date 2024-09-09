package driver

type DriverStatus uint8

// 0 = in garage, 1 = flying lap 2 = in lap, 3 = out lap, 4 = on track

const (
	InGarage DriverStatus = iota
	FlyingLap
	InLap
	OutLap
	OnTrack
)

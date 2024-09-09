package sessionType

type SessionType uint8

const (
	Unknown SessionType = iota
	P1
	P2
	P3
	ShortP
	Q1
	Q2
	Q3
	ShortQ
	OSQ
	R
	R2
	TimeTrial
)

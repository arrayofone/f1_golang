package mfd

type Mfd uint8

const (
	Closed   Mfd = 255
	CarSetup     = iota
	Pits
	Damage
	Engine
	Temperatures
)

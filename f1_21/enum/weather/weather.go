package weather

type Weather uint8

const (
	Clear Weather = iota
	LightCloud
	Overcast
	LightRain
	HeavyRain
	Storm
)

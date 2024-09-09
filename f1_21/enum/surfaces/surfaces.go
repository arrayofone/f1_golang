package surfaces

type Surface uint8

const (
	Tarmac Surface = iota
	Rumble_strip
	Concrete
	Rock
	Gravel
	Mud
	Sand
	Grass
	Water
	Cobblestone
	Metal
	Ridged
)

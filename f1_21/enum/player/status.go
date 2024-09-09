package player

type Status uint8

// 0 = not ready, 1 = ready, 2 = spectating

const (
	NotReady Status = iota
	Ready
	Spectating
)

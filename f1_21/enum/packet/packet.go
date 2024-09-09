package packet

type PacketType uint8

const (
	Motion PacketType = iota
	Session
	Lap_Data
	Event
	Participants
	Car_Setups
	Car_Telemetry
	Car_Status
	Final_Classification
	Lobby_Info
	Car_Damage
	Session_History
)

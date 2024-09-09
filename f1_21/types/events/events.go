package events

type FastestLap struct {
	VehicleIdx uint8   `json:"vehicleIdx"` // Vehicle index of car achieving fastest lap
	LapTime    float32 `json:"lapTime"`    // Lap time is in seconds
}

type Retirement struct {
	VehicleIdx uint8 `json:"vehicleIdx"` // Vehicle index of car retiring
}

type TeamMateInPits struct {
	VehicleIdx uint8 `json:"vehicleIdx"` // Vehicle index of team mate
}

type RaceWinner struct {
	VehicleIdx uint8 `json:"vehicleIdx"` // Vehicle index of the race winner
}

type PenaltyIssued struct {
	VehicleIdx       uint8 `json:"vehicleIdx"`       // Vehicle index of the car the penalty is applied to
	PenaltyType      uint8 `json:"penaltyType"`      // Penalty type – see Appendices
	InfringementType uint8 `json:"infringementType"` // Infringement type – see Appendices
	OtherVehicleIdx  uint8 `json:"otherVehicleIdx"`  // Vehicle index of the other car involved
	Time             uint8 `json:"time"`             // Time gained, or time spent doing action in seconds
	LapNum           uint8 `json:"lapNum"`           // Lap the penalty occurred on
	PlacesGained     uint8 `json:"placesGained"`     // Number of places gained by this
}

type SpeedTrapTriggered struct {
	VehicleIdx              uint8   `json:"vehicleIdx"`              // Vehicle index of the vehicle triggering speed trap
	Speed                   float32 `json:"speed"`                   // Top speed achieved in kilometres per hour
	OverallFastestInSession bool    `json:"overallFastestInSession"` // Overall fastest speed in session = 1, otherwise 0
	DriverFastestInSession  bool    `json:"driverFastestInSession"`  // Fastest speed for driver in session = 1, otherwise 0
}

type StartLights struct {
	NumLights uint8 `json:"numLights"` // Number of lights showing
}

type DriveThroughPenaltyServed struct {
	VehicleIdx uint8 `json:"vehicleIdx"` // Vehicle index of the vehicle serving drive through
}

type StopGoPenaltyServed struct {
	VehicleIdx uint8 `json:"vehicleIdx"` // Vehicle index of the vehicle serving stop go
}

type Flashback struct {
	FlashbackFrameIdentifier uint32  `json:"flashbackFrameIdentifier"` // Frame identifier flashed back to
	FlashbackSessionTime     float32 `json:"flashbackSessionTime"`     // Session time flashed back to
}

type Buttons struct {
	ButtonStatus uint32 `json:"buttonStatus"` // Bit flags specifying which buttons are being pressed currently - see appendices
}

type SessionStarted struct{}
type SessionEnded struct{}
type DRSEnabled struct{}
type DRSDisabled struct{}
type ChequeredFlag struct{}
type LightsOut struct{}

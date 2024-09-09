package packets

import (
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/enum/driver"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/enum/pit"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/enum/track"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/types/header"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/utils"
)

type LapData struct {
	LastLapTimeInMS             uint32              `json:"lastLapTimeInMS"`             // Last lap time in milliseconds
	CurrentLapTimeInMS          uint32              `json:"currentLapTimeInMS"`          // Current time around the lap in milliseconds
	Sector1TimeInMS             uint16              `json:"sector1TimeInMS"`             // Sector 1 time in milliseconds
	Sector2TimeInMS             uint16              `json:"sector2TimeInMS"`             // Sector 2 time in milliseconds
	LapDistance                 float32             `json:"lapDistance"`                 // Distance vehicle is around current lap in metres – could be negative if line hasn’t been crossed yet
	TotalDistance               float32             `json:"totalDistance"`               // Total distance travelled in session in metres – could be negative if line hasn’t been crossed yet
	SafetyCarDelta              float32             `json:"safetyCarDelta"`              // Delta in seconds for safety car
	CarPosition                 uint8               `json:"carPosition"`                 // Car race position
	CurrentLapNum               uint8               `json:"currentLapNum"`               // Current lap number
	PitStatus                   pit.PitStatus       `json:"pitStatus"`                   // 0 = none, 1 = pitting, 2 = in pit area
	NumPitStops                 uint8               `json:"numPitStops"`                 // Number of pit stops taken in this race
	Sector                      track.Sector        `json:"sector"`                      // 0 = sector1, 1 = sector2, 2 = sector3
	CurrentLapInvalid           bool                `json:"currentLapInvalid"`           // Current lap invalid - 0 = valid, 1 = invalid
	Penalties                   uint8               `json:"penalties"`                   // Accumulated time penalties in seconds to be added
	Warnings                    uint8               `json:"warnings"`                    // Accumulated number of warnings issued
	NumUnservedDriveThroughPens uint8               `json:"numUnservedDriveThroughPens"` // Num drive through pens left to serve
	NumUnservedStopGoPens       uint8               `json:"numUnservedStopGoPens"`       // Num stop go pens left to serve
	GridPosition                uint8               `json:"gridPosition"`                // Grid position the vehicle started the race in
	DriverStatus                driver.DriverStatus `json:"driverStatus"`                // Status of driver - 0 = in garage, 1 = flying lap 2 = in lap, 3 = out lap, 4 = on track
	ResultStatus                driver.ResultStatus `json:"resultStatus"`                // Result status - 0 = invalid, 1 = inactive, 2 = active 3 = finished, 4 = didnotfinish, 5 = disqualified 6 = not classified, 7 = retired
	PitLaneTimerActive          uint8               `json:"pitLaneTimerActive"`          // Pit lane timing, 0 = inactive, 1 = active
	PitLaneTimeInLaneInMS       uint16              `json:"pitLaneTimeInLaneInMS"`       // If active, the current time spent in the pit lane in ms
	PitStopTimerInMS            uint16              `json:"pitStopTimerInMS"`            // Time of the actual pit stop in ms
	PitStopShouldServePen       uint8               `json:"pitStopShouldServePen"`       // Whether the car should serve a penalty at this stop
}

func newLapData(bc *utils.ByteCursor) LapData {
	return LapData{
		LastLapTimeInMS:             bc.Uint32(),
		CurrentLapTimeInMS:          bc.Uint32(),
		Sector1TimeInMS:             bc.Uint16(),
		Sector2TimeInMS:             bc.Uint16(),
		LapDistance:                 bc.Float(),
		TotalDistance:               bc.Float(),
		SafetyCarDelta:              bc.Float(),
		CarPosition:                 bc.Uint8(),
		CurrentLapNum:               bc.Uint8(),
		PitStatus:                   pit.PitStatus(bc.Uint8()),
		NumPitStops:                 bc.Uint8(),
		Sector:                      track.Sector(bc.Uint8()),
		CurrentLapInvalid:           bc.Bool(),
		Penalties:                   bc.Uint8(),
		Warnings:                    bc.Uint8(),
		NumUnservedDriveThroughPens: bc.Uint8(),
		NumUnservedStopGoPens:       bc.Uint8(),
		GridPosition:                bc.Uint8(),
		DriverStatus:                driver.DriverStatus(bc.Uint8()),
		ResultStatus:                driver.ResultStatus(bc.Uint8()),
		PitLaneTimerActive:          bc.Uint8(),
		PitLaneTimeInLaneInMS:       bc.Uint16(),
		PitStopTimerInMS:            bc.Uint16(),
		PitStopShouldServePen:       bc.Uint8(),
	}
}

type PacketLapData struct {
	Header  header.Header `json:"header"`  // Header
	LapData [22]LapData   `json:"lapData"` // Lap data for all cars on track
}

func NewPacketLapData(b []byte) *PacketLapData {
	bc := utils.NewByteCursor(b)

	header := header.DecodeHeader(bc.B(24))
	LapData := [22]LapData{}

	for i := 0; i < 22; i++ {
		LapData[i] = newLapData(bc)
	}

	return &PacketLapData{
		Header:  header,
		LapData: LapData,
	}
}

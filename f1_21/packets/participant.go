package packets

import (
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/enum/nationality"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/enum/teams"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/enum/telemetry"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/types/header"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/utils"
)

type ParticipantData struct {
	AiControlled  bool                      `json:"aiControlled"`  // Whether the vehicle is AI (1) or Human (0) controlled
	DriverId      uint8                     `json:"driverId"`      // Driver id - see appendix, 255 if network human
	NetworkId     uint8                     `json:"networkId"`     // Network id – unique identifier for network players
	TeamId        teams.Team                `json:"teamId"`        // Team id - see appendix
	MyTeam        uint8                     `json:"myTeam"`        // My team flag – 1 = My Team, 0 = otherwise
	RaceNumber    uint8                     `json:"raceNumber"`    // Race number of the car
	Nationality   nationality.Nationality   `json:"nationality"`   // Nationality of the driver
	Name          string                    `json:"name"`          // Name of participant in UTF-8 format – null terminated Will be truncated with … (U+2026) if too long
	YourTelemetry telemetry.TelemetryStatus `json:"yourTelemetry"` // The player's UDP setting, 0 = restricted, 1 = public
}

func newParticipantData(bc *utils.ByteCursor) ParticipantData {
	return ParticipantData{
		AiControlled:  bc.Bool(),
		DriverId:      bc.Uint8(),
		NetworkId:     bc.Uint8(),
		TeamId:        teams.Team(bc.Uint8()),
		MyTeam:        bc.Uint8(),
		RaceNumber:    bc.Uint8(),
		Nationality:   nationality.Nationality(bc.Uint8()),
		Name:          bc.String(48),
		YourTelemetry: telemetry.TelemetryStatus(bc.Uint8()),
	}
}

type PacketParticipantsData struct {
	Header header.Header `json:"header"` // Header

	NumActiveCars uint8             `json:"numActiveCars"` // Number of active cars in the data – should match number of cars on HUD
	Participants  []ParticipantData `json:"participants"`
}

func NewPacketParticipantsData(b []byte) *PacketParticipantsData {
	bc := utils.NewByteCursor(b)

	header := header.DecodeHeader(bc.B(24))
	numActiveCars := bc.Uint8()
	participants := make([]ParticipantData, numActiveCars)

	for i := 0; i < int(numActiveCars); i++ {
		participants[i] = newParticipantData(bc)
	}

	return &PacketParticipantsData{
		Header:        header,
		NumActiveCars: numActiveCars,
		Participants:  participants,
	}
}

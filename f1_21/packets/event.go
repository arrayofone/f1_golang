package packets

import (
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/enum/event"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/types/events"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/types/header"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/utils"
)

type EventDataDetails interface{}
type PacketEventData struct {
	Header header.Header `json:"header"` // Header

	EventStringCode event.EventCode  `json:"eventStringCode"` // Event string code, see below
	EventDetails    EventDataDetails `json:"eventDetails"`    // Event details - should be interpreted differently for each type
}

func NewPacketEventData(b []byte) *PacketEventData {
	bc := utils.NewByteCursor(b)

	header := header.DecodeHeader(bc.B(24))
	strCode := bc.String(4)

	var EventDetails EventDataDetails

	switch event.EventCode(strCode) {
	case event.FastestLap:
		EventDetails = events.FastestLap{
			VehicleIdx: bc.Uint8(),
			LapTime:    bc.Float(),
		}
	case event.Retirement:
		EventDetails = events.Retirement{
			VehicleIdx: bc.Uint8(),
		}
	case event.TeammateInPits:
		EventDetails = events.TeamMateInPits{
			VehicleIdx: bc.Uint8(),
		}
	case event.RaceWinner:
		EventDetails = events.RaceWinner{
			VehicleIdx: bc.Uint8(),
		}
	case event.StartLights:
		EventDetails = events.StartLights{
			NumLights: bc.Uint8(),
		}
	case event.Flashback:
		EventDetails = events.Flashback{
			FlashbackFrameIdentifier: bc.Uint32(),
			FlashbackSessionTime:     bc.Float(),
		}
	case event.ButtonStatus:
		EventDetails = events.Buttons{
			ButtonStatus: bc.Uint32(),
		}

	case event.SpeedTrapTriggered:
		EventDetails = events.SpeedTrapTriggered{
			VehicleIdx:              bc.Uint8(),
			Speed:                   bc.Float(),
			OverallFastestInSession: bc.Bool(),
			DriverFastestInSession:  bc.Bool(),
		}

	case event.PenaltyIssued:
		EventDetails = events.PenaltyIssued{
			PenaltyType:      bc.Uint8(),
			InfringementType: bc.Uint8(),
			VehicleIdx:       bc.Uint8(),
			OtherVehicleIdx:  bc.Uint8(),
			Time:             bc.Uint8(),
			LapNum:           bc.Uint8(),
			PlacesGained:     bc.Uint8(),
		}

	case event.DriveThroughPenaltyServed:
		EventDetails = events.DriveThroughPenaltyServed{
			VehicleIdx: bc.Uint8(),
		}

	case event.StopGoPenaltyServed:
		EventDetails = events.StopGoPenaltyServed{
			VehicleIdx: bc.Uint8(),
		}

	//these don't have extra data but they are created for type checking for the receiver
	case event.SessionStarted:
		EventDetails = events.SessionStarted{}
	case event.SessionEnded:
		EventDetails = events.SessionEnded{}
	case event.DRSEnabled:
		EventDetails = events.DRSEnabled{}
	case event.DRSDisabled:
		EventDetails = events.DRSDisabled{}
	case event.ChequeredFlag:
		EventDetails = events.ChequeredFlag{}
	case event.LightsOut:
		EventDetails = events.LightsOut{}
	}

	return &PacketEventData{
		Header:          header,
		EventStringCode: event.EventCode(strCode),
		EventDetails:    EventDetails,
	}
}

package f1_21

import (
	"errors"

	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/enum/packet"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/packets"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/utils"
)

type F121 struct {
	Config *Config
}

func NewF121(c *Config) *F121 {
	return &F121{
		Config: c,
	}
}

func (f *F121) Write(b []byte) (int, error) {
	//the 6th byte defines the packet type
	switch packet.PacketType(utils.D_uint8(b[5:6])) {
	case packet.Motion:
		f.Config.OnPacketCarMotionData(packets.NewPacketCarMotionData(b))
	case packet.Session:
		f.Config.OnPacketSessionData(packets.NewPacketSessionData(b))
	case packet.Lap_Data:
		f.Config.OnPacketLapData(packets.NewPacketLapData(b))
	case packet.Event:
		f.Config.OnPacketEventData(packets.NewPacketEventData(b))
	case packet.Participants:
		f.Config.OnPacketParticipantsData(packets.NewPacketParticipantsData(b))
	case packet.Car_Setups:
		f.Config.OnPacketCarSetupData(packets.NewPacketCarSetupData(b))
	case packet.Car_Telemetry:
		f.Config.OnPacketCarTelemetryData(packets.NewPacketCarTelemetryData(b))
	case packet.Car_Status:
		f.Config.OnPacketCarStatusData(packets.NewPacketCarStatusData(b))
	case packet.Final_Classification:
		f.Config.OnPacketFinalClassificationData(packets.NewPacketFinalClassificationData(b))
	case packet.Lobby_Info:
		f.Config.OnPacketLobbyInfoData(packets.NewPacketLobbyInfoData(b))
	case packet.Car_Damage:
		f.Config.OnPacketCarDamageData(packets.NewPacketCarDamageData(b))
	case packet.Session_History:
		f.Config.OnPacketSessionHistoryData(packets.NewPacketSessionHistoryData(b))
	default:
		return -1, errors.New("unknown packet type")
	}

	return len(b), nil
}

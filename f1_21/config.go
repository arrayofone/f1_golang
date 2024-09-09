package f1_21

import "gitlab.com/DarrenBangsund/cmf1_golang/f1_21/packets"

type Config struct {
	OnPacketCarMotionData           func(*packets.PacketCarMotionData)
	OnPacketSessionData             func(*packets.PacketSessionData)
	OnPacketLapData                 func(*packets.PacketLapData)
	OnPacketEventData               func(*packets.PacketEventData)
	OnPacketParticipantsData        func(*packets.PacketParticipantsData)
	OnPacketCarSetupData            func(*packets.PacketCarSetupData)
	OnPacketCarTelemetryData        func(*packets.PacketCarTelemetryData)
	OnPacketCarStatusData           func(*packets.PacketCarStatusData)
	OnPacketFinalClassificationData func(*packets.PacketFinalClassificationData)
	OnPacketLobbyInfoData           func(*packets.PacketLobbyInfoData)
	OnPacketCarDamageData           func(*packets.PacketCarDamageData)
	OnPacketSessionHistoryData      func(*packets.PacketSessionHistoryData)
}

func DefaultConfig() *Config {
	return &Config{
		OnPacketCarMotionData:           func(p *packets.PacketCarMotionData) {},
		OnPacketSessionData:             func(p *packets.PacketSessionData) {},
		OnPacketLapData:                 func(p *packets.PacketLapData) {},
		OnPacketEventData:               func(p *packets.PacketEventData) {},
		OnPacketParticipantsData:        func(p *packets.PacketParticipantsData) {},
		OnPacketCarSetupData:            func(p *packets.PacketCarSetupData) {},
		OnPacketCarTelemetryData:        func(p *packets.PacketCarTelemetryData) {},
		OnPacketCarStatusData:           func(p *packets.PacketCarStatusData) {},
		OnPacketFinalClassificationData: func(p *packets.PacketFinalClassificationData) {},
		OnPacketLobbyInfoData:           func(p *packets.PacketLobbyInfoData) {},
		OnPacketCarDamageData:           func(p *packets.PacketCarDamageData) {},
		OnPacketSessionHistoryData:      func(p *packets.PacketSessionHistoryData) {},
	}
}

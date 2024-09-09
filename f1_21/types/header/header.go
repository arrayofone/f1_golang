package header

import (
	"fmt"

	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/enum/packet"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/utils"
)

type Header struct {
	PacketFormat            uint16            `json:"packetFormat"`
	GameMajorVersion        uint8             `json:"gameMajorVersion"`
	GameMinorVersion        uint8             `json:"gameMinorVersion"`
	PacketVersion           uint8             `json:"packetVersion"`
	PacketId                packet.PacketType `json:"packetId"`
	SessionUID              string            `json:"sessionUID"`
	SessionTime             float32           `json:"sessionTime"`
	FrameIdentifier         uint32            `json:"frameIdentifier"`
	PlayerCarIndex          uint8             `json:"playerCarIndex"`
	SecondaryPlayerCarIndex uint8             `json:"secondaryPlayerCarIndex"`
}

func DecodeHeader(b []byte) Header {
	_ = b[23]
	bc := utils.NewByteCursor(b)
	return Header{
		PacketFormat:            bc.Uint16(),
		GameMajorVersion:        bc.Uint8(),
		GameMinorVersion:        bc.Uint8(),
		PacketVersion:           bc.Uint8(),
		PacketId:                packet.PacketType(bc.Uint8()),
		SessionUID:              fmt.Sprintf("%d", bc.Uint64()),
		SessionTime:             bc.Float(),
		FrameIdentifier:         bc.Uint32(),
		PlayerCarIndex:          bc.Uint8(),
		SecondaryPlayerCarIndex: bc.Uint8(),
	}
}

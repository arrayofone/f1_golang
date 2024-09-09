package packets

import (
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/enum/nationality"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/enum/player"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/enum/teams"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/types/header"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/utils"
)

type LobbyInfoData struct {
	AiControlled uint8                   `json:"aiControlled"` // Whether the vehicle is AI (1) or Human (0) controlled
	TeamId       teams.Team              `json:"teamId"`       // Team id - see appendix (255 if no team currently selected)
	Nationality  nationality.Nationality `json:"nationality"`  // Nationality of the driver
	Name         string                  `json:"name"`         // Name of participant in UTF-8 format – null terminated Will be truncated with ... (U+2026) if too long
	CarNumber    uint8                   `json:"carNumber"`    // Car number of the player
	ReadyStatus  player.Status           `json:"readyStatus"`  // 0 = not ready, 1 = ready, 2 = spectating
}

func newLobbyInfoData(bc *utils.ByteCursor) LobbyInfoData {
	return LobbyInfoData{
		AiControlled: bc.Uint8(),
		TeamId:       teams.Team(bc.Uint8()),
		Nationality:  nationality.Nationality(bc.Uint8()),
		Name:         bc.String(48),
		CarNumber:    bc.Uint8(),
		ReadyStatus:  player.Status(bc.Uint8()),
	}
}

type PacketLobbyInfoData struct {
	Header header.Header `json:"header"` // Header

	NumPlayers   uint8           `json:"numPlayers"` // Number of players in the lobby data
	LobbyPlayers []LobbyInfoData `json:"lobbyPlayers"`
}

func NewPacketLobbyInfoData(b []byte) *PacketLobbyInfoData {
	bc := utils.NewByteCursor(b)

	header := header.DecodeHeader(bc.B(24))
	numPlayers := bc.Uint8()
	lobbyPlayers := make([]LobbyInfoData, numPlayers)

	for i := 0; uint8(i) < numPlayers; i++ {
		lobbyPlayers[i] = newLobbyInfoData(bc)
	}

	return &PacketLobbyInfoData{
		Header:       header,
		NumPlayers:   numPlayers,
		LobbyPlayers: lobbyPlayers,
	}
}

package packets

import (
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/types/header"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/utils"
)

type LapHistoryData struct {
	LapTimeInMS      uint32 `json:"lapTimeInMS"`      // Lap time in milliseconds
	Sector1TimeInMS  uint16 `json:"sector1TimeInMS"`  // Sector 1 time in milliseconds
	Sector2TimeInMS  uint16 `json:"sector2TimeInMS"`  // Sector 2 time in milliseconds
	Sector3TimeInMS  uint16 `json:"sector3TimeInMS"`  // Sector 3 time in milliseconds
	LapValidBitFlags uint8  `json:"lapValidBitFlags"` // 0x01 bit set-lap valid,      0x02 bit set-sector 1 valid 0x04 bit set-sector 2 valid, 0x08 bit set-sector 3 valid
}

func newLapHistoryData(bc *utils.ByteCursor) LapHistoryData {
	return LapHistoryData{
		LapTimeInMS:      bc.Uint32(),
		Sector1TimeInMS:  bc.Uint16(),
		Sector2TimeInMS:  bc.Uint16(),
		Sector3TimeInMS:  bc.Uint16(),
		LapValidBitFlags: bc.Uint8(),
	}
}

type TyreStintHistoryData struct {
	EndLap             uint8 `json:"endLap"`             // Lap the tyre usage ends on (255 of current tyre)
	TyreActualCompound uint8 `json:"tyreActualCompound"` // Actual tyres used by this driver
	TyreVisualCompound uint8 `json:"tyreVisualCompound"` // Visual tyres used by this driver
}

func newTyreStintHistoryData(bc *utils.ByteCursor) TyreStintHistoryData {
	return TyreStintHistoryData{
		EndLap:             bc.Uint8(),
		TyreActualCompound: bc.Uint8(),
		TyreVisualCompound: bc.Uint8(),
	}
}

type PacketSessionHistoryData struct {
	Header                header.Header          `json:"header"`            // Header
	CarIdx                uint8                  `json:"carIdx"`            // Index of the car this lap data relates to
	NumLaps               uint8                  `json:"numLaps"`           // Num laps in the data (including current partial lap)
	NumTyreStints         uint8                  `json:"numTyreStints"`     // Number of tyre stints in the data
	BestLapTimeLapNum     uint8                  `json:"bestLapTimeLapNum"` // Lap the best lap time was achieved on
	BestSector1LapNum     uint8                  `json:"bestSector1LapNum"` // Lap the best Sector 1 time was achieved on
	BestSector2LapNum     uint8                  `json:"bestSector2LapNum"` // Lap the best Sector 2 time was achieved on
	BestSector3LapNum     uint8                  `json:"bestSector3LapNum"` // Lap the best Sector 3 time was achieved on
	LapHistoryData        []LapHistoryData       `json:"lapHistoryData"`    // 100 laps of data max
	TyreStintsHistoryData []TyreStintHistoryData `json:"tyreStintsHistoryData"`
}

func NewPacketSessionHistoryData(b []byte) *PacketSessionHistoryData {
	bc := utils.NewByteCursor(b)

	header := header.DecodeHeader(bc.B(24))
	carIdx := bc.Uint8()
	numLaps := bc.Uint8()
	numTyreStints := bc.Uint8()
	bestLapTimeLapNum := bc.Uint8()
	bestSector1LapNum := bc.Uint8()
	bestSector2LapNum := bc.Uint8()
	bestSector3LapNum := bc.Uint8()
	lapHistoryData := make([]LapHistoryData, numLaps)
	tyreStintsHistoryData := make([]TyreStintHistoryData, numTyreStints)

	for i := 0; uint8(i) < numLaps; i++ {
		lapHistoryData[i] = newLapHistoryData(bc)
	}

	for i := 0; uint8(i) < numTyreStints; i++ {
		tyreStintsHistoryData[i] = newTyreStintHistoryData(bc)
	}

	return &PacketSessionHistoryData{
		Header:                header,
		CarIdx:                carIdx,
		NumLaps:               numLaps,
		NumTyreStints:         numTyreStints,
		BestLapTimeLapNum:     bestLapTimeLapNum,
		BestSector1LapNum:     bestSector1LapNum,
		BestSector2LapNum:     bestSector2LapNum,
		BestSector3LapNum:     bestSector3LapNum,
		LapHistoryData:        lapHistoryData,
		TyreStintsHistoryData: tyreStintsHistoryData,
	}
}

package packets

import (
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/enum/tyres"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/types/header"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/utils"
)

type FinalClassificationData struct {
	Position         uint8             `json:"position"`         // Finishing position
	NumLaps          uint8             `json:"numLaps"`          // Number of laps completed
	GridPosition     uint8             `json:"gridPosition"`     // Grid position of the car
	Points           uint8             `json:"points"`           // Number of points scored
	NumPitStops      uint8             `json:"numPitStops"`      // Number of pit stops made
	ResultStatus     uint8             `json:"resultStatus"`     // Result status - 0 = invalid, 1 = inactive, 2 = active 3 = finished, 4 = didnotfinish, 5 = disqualified 6 = not classified, 7 = retired
	BestLapTimeInMS  uint32            `json:"bestLapTimeInMS"`  // Best lap time of the session in milliseconds
	TotalRaceTime    uint64            `json:"totalRaceTime"`    // Total race time in seconds without penalties
	PenaltiesTime    uint8             `json:"penaltiesTime"`    // Total penalties accumulated in seconds
	NumPenalties     uint8             `json:"numPenalties"`     // Number of penalties applied to this driver
	NumTyreStints    uint8             `json:"numTyreStints"`    // Number of tyres stints up to maximum
	TyreStintsActual [8]tyres.Compound `json:"tyreStintsActual"` // Actual tyres used by this driver
	TyreStintsVisual [8]tyres.Visual   `json:"tyreStintsVisual"` // Visual tyres used by this driver
}

func newFinalClassificationData(bc *utils.ByteCursor) FinalClassificationData {
	f := FinalClassificationData{
		Position:        bc.Uint8(),
		NumLaps:         bc.Uint8(),
		GridPosition:    bc.Uint8(),
		Points:          bc.Uint8(),
		NumPitStops:     bc.Uint8(),
		ResultStatus:    bc.Uint8(),
		BestLapTimeInMS: bc.Uint32(),
		TotalRaceTime:   bc.Uint64(),
		PenaltiesTime:   bc.Uint8(),
		NumPenalties:    bc.Uint8(),
		NumTyreStints:   bc.Uint8(),
		// TyreStintsActual: ...,
		// TyreStintsVisual: ...,
	}

	tsa := [8]tyres.Compound{}
	for i := 0; i < 8; i++ {
		tsa[i] = tyres.Compound(bc.Uint8())
	}

	tsv := [8]tyres.Visual{}
	for i := 0; i < 8; i++ {
		tsv[i] = tyres.Visual(bc.Uint8())
	}

	f.TyreStintsActual = tsa
	f.TyreStintsVisual = tsv

	return f
}

type PacketFinalClassificationData struct {
	Header             header.Header             `json:"header"`  // Header
	NumCars            uint8                     `json:"numCars"` // Number of cars in the final classification
	ClassificationData []FinalClassificationData `json:"classificationData"`
}

func NewPacketFinalClassificationData(b []byte) *PacketFinalClassificationData {
	bc := utils.NewByteCursor(b)

	header := header.DecodeHeader(bc.B(24))
	numCars := bc.Uint8()
	classificationData := make([]FinalClassificationData, numCars)

	for i := 0; uint8(i) < numCars; i++ {
		classificationData[i] = newFinalClassificationData(bc)
	}

	return &PacketFinalClassificationData{
		Header:             header,
		NumCars:            numCars,
		ClassificationData: classificationData,
	}
}

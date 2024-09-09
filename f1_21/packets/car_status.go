package packets

import (
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/enum/flag"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/enum/tyres"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/types/header"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/utils"
)

type CarStatusData struct {
	TractionControl         uint8          `json:"tractionControl"`         // Traction control - 0 = off, 1 = medium, 2 = full
	AntiLockBrakes          bool           `json:"antiLockBrakes"`          // 0 (off) - 1 (on)
	FuelMix                 uint8          `json:"fuelMix"`                 // Fuel mix - 0 = lean, 1 = standard, 2 = rich, 3 = max
	FrontBrakeBias          uint8          `json:"frontBrakeBias"`          // Front brake bias (percentage)
	PitLimiterStatus        bool           `json:"pitLimiterStatus"`        // Pit limiter status - 0 = off, 1 = on
	FuelInTank              float32        `json:"fuelInTank"`              // Current fuel mass
	FuelCapacity            float32        `json:"fuelCapacity"`            // Fuel capacity
	FuelRemainingLaps       float32        `json:"fuelRemainingLaps"`       // Fuel remaining in terms of laps (value on MFD)
	MaxRPM                  uint16         `json:"maxRPM"`                  // Cars max RPM, point of rev limiter
	IdleRPM                 uint16         `json:"idleRPM"`                 // Cars idle RPM
	MaxGears                uint8          `json:"maxGears"`                // Maximum number of gears
	DrsAllowed              bool           `json:"drsAllowed"`              // 0 = not allowed, 1 = allowed
	DrsActivationDistance   uint16         `json:"drsActivationDistance"`   // 0 = DRS not available, non-zero - DRS will be available in [X] metres
	ActualTyreCompound      tyres.Compound `json:"actualTyreCompound"`      // F1 Modern - 16 = C5, 17 = C4, 18 = C3, 19 = C2, 20 = C1 7 = inter, 8 = wet F1 Classic - 9 = dry, 10 = wet F2 – 11 = super soft, 12 = soft, 13 = medium, 14 = hard 15 = wet
	VisualTyreCompound      tyres.Visual   `json:"visualTyreCompound"`      // F1 visual (can be different from actual compound) 16 = soft, 17 = medium, 18 = hard, 7 = inter, 8 = wet F1 Classic – same as above F2 ‘19, 15 = wet, 19 – super soft, 20 = soft 21 = medium , 22 = hard
	TyresAgeLaps            uint8          `json:"tyresAgeLaps"`            // Age in laps of the current set of tyres
	VehicleFiaFlags         flag.Flag      `json:"vehicleFiaFlags"`         // -1 = invalid/unknown, 0 = none, 1 = green 2 = blue, 3 = yellow, 4 = red
	ErsStoreEnergy          float32        `json:"ersStoreEnergy"`          // ERS energy store in Joules
	ErsDeployMode           uint8          `json:"ersDeployMode"`           // ERS deployment mode, 0 = none, 1 = medium 2 = hotlap, 3 = overtake
	ErsHarvestedThisLapMGUK float32        `json:"ersHarvestedThisLapMGUK"` // ERS energy harvested this lap by MGU-K
	ErsHarvestedThisLapMGUH float32        `json:"ersHarvestedThisLapMGUH"` // ERS energy harvested this lap by MGU-H
	ErsDeployedThisLap      float32        `json:"ersDeployedThisLap"`      // ERS energy deployed this lap
	NetworkPaused           bool           `json:"networkPaused"`           // Whether the car is paused in a network game
}

func newCarStatusData(bc *utils.ByteCursor) CarStatusData {
	return CarStatusData{
		TractionControl:         bc.Uint8(),
		AntiLockBrakes:          bc.Bool(),
		FuelMix:                 bc.Uint8(),
		FrontBrakeBias:          bc.Uint8(),
		PitLimiterStatus:        bc.Bool(),
		FuelInTank:              bc.Float(),
		FuelCapacity:            bc.Float(),
		FuelRemainingLaps:       bc.Float(),
		MaxRPM:                  bc.Uint16(),
		IdleRPM:                 bc.Uint16(),
		MaxGears:                bc.Uint8(),
		DrsAllowed:              bc.Bool(),
		DrsActivationDistance:   bc.Uint16(),
		ActualTyreCompound:      tyres.Compound(bc.Uint8()),
		VisualTyreCompound:      tyres.Visual(bc.Uint8()),
		TyresAgeLaps:            bc.Uint8(),
		VehicleFiaFlags:         flag.Flag(bc.Uint8()),
		ErsStoreEnergy:          bc.Float(),
		ErsDeployMode:           bc.Uint8(),
		ErsHarvestedThisLapMGUK: bc.Float(),
		ErsHarvestedThisLapMGUH: bc.Float(),
		ErsDeployedThisLap:      bc.Float(),
		NetworkPaused:           bc.Bool(),
	}
}

type PacketCarStatusData struct {
	Header header.Header `json:"header"` // Header

	CarStatusData [22]CarStatusData `json:"carStatusData"`
}

func NewPacketCarStatusData(b []byte) *PacketCarStatusData {
	bc := utils.NewByteCursor(b)

	header := header.DecodeHeader(bc.B(24))
	CarStatusData := [22]CarStatusData{}

	for i := 0; i < 22; i++ {
		CarStatusData[i] = newCarStatusData(bc)
	}

	return &PacketCarStatusData{
		Header:        header,
		CarStatusData: CarStatusData,
	}
}

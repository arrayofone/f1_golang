package packets

import (
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/types/header"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/utils"
)

type CarSetupData struct {
	FrontWing              uint8   `json:"frontWing"`              // Front wing aero
	RearWing               uint8   `json:"rearWing"`               // Rear wing aero
	OnThrottle             uint8   `json:"onThrottle"`             // Differential adjustment on throttle (percentage)
	OffThrottle            uint8   `json:"offThrottle"`            // Differential adjustment off throttle (percentage)
	FrontCamber            float32 `json:"frontCamber"`            // Front camber angle (suspension geometry)
	RearCamber             float32 `json:"rearCamber"`             // Rear camber angle (suspension geometry)
	FrontToe               float32 `json:"frontToe"`               // Front toe angle (suspension geometry)
	RearToe                float32 `json:"rearToe"`                // Rear toe angle (suspension geometry)
	FrontSuspension        uint8   `json:"frontSuspension"`        // Front suspension
	RearSuspension         uint8   `json:"rearSuspension"`         // Rear suspension
	FrontAntiRollBar       uint8   `json:"frontAntiRollBar"`       // Front anti-roll bar
	RearAntiRollBar        uint8   `json:"rearAntiRollBar"`        // Front anti-roll bar
	FrontSuspensionHeight  uint8   `json:"frontSuspensionHeight"`  // Front ride height
	RearSuspensionHeight   uint8   `json:"rearSuspensionHeight"`   // Rear ride height
	BrakePressure          uint8   `json:"brakePressure"`          // Brake pressure (percentage)
	BrakeBias              uint8   `json:"brakeBias"`              // Brake bias (percentage)
	RearLeftTyrePressure   float32 `json:"rearLeftTyrePressure"`   // Rear left tyre pressure (PSI)
	RearRightTyrePressure  float32 `json:"rearRightTyrePressure"`  // Rear right tyre pressure (PSI)
	FrontLeftTyrePressure  float32 `json:"frontLeftTyrePressure"`  // Front left tyre pressure (PSI)
	FrontRightTyrePressure float32 `json:"frontRightTyrePressure"` // Front right tyre pressure (PSI)
	Ballast                uint8   `json:"ballast"`                // Ballast
	FuelLoad               float32 `json:"fuelLoad"`               // Fuel load
}

func newCarSetupData(bc *utils.ByteCursor) CarSetupData {
	return CarSetupData{
		FrontWing:              bc.Uint8(),
		RearWing:               bc.Uint8(),
		OnThrottle:             bc.Uint8(),
		OffThrottle:            bc.Uint8(),
		FrontCamber:            bc.Float(),
		RearCamber:             bc.Float(),
		FrontToe:               bc.Float(),
		RearToe:                bc.Float(),
		FrontSuspension:        bc.Uint8(),
		RearSuspension:         bc.Uint8(),
		FrontAntiRollBar:       bc.Uint8(),
		RearAntiRollBar:        bc.Uint8(),
		FrontSuspensionHeight:  bc.Uint8(),
		RearSuspensionHeight:   bc.Uint8(),
		BrakePressure:          bc.Uint8(),
		BrakeBias:              bc.Uint8(),
		RearLeftTyrePressure:   bc.Float(),
		RearRightTyrePressure:  bc.Float(),
		FrontLeftTyrePressure:  bc.Float(),
		FrontRightTyrePressure: bc.Float(),
		Ballast:                bc.Uint8(),
		FuelLoad:               bc.Float(),
	}
}

type PacketCarSetupData struct {
	Header header.Header `json:"header"` // Header

	CarSetups [22]CarSetupData `json:"carSetups"` //length of 22
}

func NewPacketCarSetupData(b []byte) *PacketCarSetupData {
	bc := utils.NewByteCursor(b)

	header := header.DecodeHeader(bc.B(24))
	carSetups := [22]CarSetupData{}

	for i := 0; i < 22; i++ {
		carSetups[i] = newCarSetupData(bc)
	}

	return &PacketCarSetupData{
		Header:    header,
		CarSetups: carSetups,
	}
}

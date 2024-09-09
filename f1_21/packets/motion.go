package packets

import (
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/types/header"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/types/wheels"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/utils"
)

type CarMotionData struct { // 60 bytes
	WorldPositionX     float32 `json:"worldPositionX"`     // World space X position
	WorldPositionY     float32 `json:"worldPositionY"`     // World space Y position
	WorldPositionZ     float32 `json:"worldPositionZ"`     // World space Z position
	WorldVelocityX     float32 `json:"worldVelocityX"`     // Velocity in world space X
	WorldVelocityY     float32 `json:"worldVelocityY"`     // Velocity in world space Y
	WorldVelocityZ     float32 `json:"worldVelocityZ"`     // Velocity in world space Z
	WorldForwardDirX   float32 `json:"worldForwardDirX"`   // World space forward X direction (normalised)
	WorldForwardDirY   float32 `json:"worldForwardDirY"`   // World space forward Y direction (normalised)
	WorldForwardDirZ   float32 `json:"worldForwardDirZ"`   // World space forward Z direction (normalised)
	WorldRightDirX     float32 `json:"worldRightDirX"`     // World space right X direction (normalised)
	WorldRightDirY     float32 `json:"worldRightDirY"`     // World space right Y direction (normalised)
	WorldRightDirZ     float32 `json:"worldRightDirZ"`     // World space right Z direction (normalised)
	GForceLateral      float32 `json:"gForceLateral"`      // Lateral G-Force component
	GForceLongitudinal float32 `json:"gForceLongitudinal"` // Longitudinal G-Force component
	GForceVertical     float32 `json:"gForceVertical"`     // Vertical G-Force component
	Yaw                float32 `json:"yaw"`                // Yaw angle in radians
	Pitch              float32 `json:"pitch"`              // Pitch angle in radians
	Roll               float32 `json:"roll"`               // Roll angle in radians
}

func newCarMotionData(bc *utils.ByteCursor) CarMotionData {
	return CarMotionData{
		WorldPositionX:     bc.Float(),
		WorldPositionY:     bc.Float(),
		WorldPositionZ:     bc.Float(),
		WorldVelocityX:     bc.Float(),
		WorldVelocityY:     bc.Float(),
		WorldVelocityZ:     bc.Float(),
		WorldForwardDirX:   float32(bc.Int16() / 32767.0),
		WorldForwardDirY:   float32(bc.Int16() / 32767.0),
		WorldForwardDirZ:   float32(bc.Int16() / 32767.0),
		WorldRightDirX:     float32(bc.Int16() / 32767.0),
		WorldRightDirY:     float32(bc.Int16() / 32767.0),
		WorldRightDirZ:     float32(bc.Int16() / 32767.0),
		GForceLateral:      bc.Float(),
		GForceLongitudinal: bc.Float(),
		GForceVertical:     bc.Float(),
		Yaw:                bc.Float(),
		Pitch:              bc.Float(),
		Roll:               bc.Float(),
	}
}

type PacketCarMotionData struct {
	Header header.Header `json:"header"`

	CarMotionData [22]CarMotionData `json:"carMotionData"` // Data for all cars on track

	// Extra player car ONLY data
	SuspensionPosition     wheels.WheelsF `json:"suspensionPosition"`     // Note: All wheel arrays have the following order:
	SuspensionVelocity     wheels.WheelsF `json:"suspensionVelocity"`     // RL, RR, FL, FR
	SuspensionAcceleration wheels.WheelsF `json:"suspensionAcceleration"` // RL, RR, FL, FR
	WheelSpeed             wheels.WheelsF `json:"wheelSpeed"`             // Speed of each wheel
	WheelSlip              wheels.WheelsF `json:"wheelSlip"`              // Slip ratio for each wheel
	LocalVelocityX         float32        `json:"localVelocityX"`         // Velocity in local space
	LocalVelocityY         float32        `json:"localVelocityY"`         // Velocity in local space
	LocalVelocityZ         float32        `json:"localVelocityZ"`         // Velocity in local space
	AngularVelocityX       float32        `json:"angularVelocityX"`       // Angular velocity x-component
	AngularVelocityY       float32        `json:"angularVelocityY"`       // Angular velocity y-component
	AngularVelocityZ       float32        `json:"angularVelocityZ"`       // Angular velocity z-component
	AngularAccelerationX   float32        `json:"angularAccelerationX"`   // Angular velocity x-component
	AngularAccelerationY   float32        `json:"angularAccelerationY"`   // Angular velocity y-component
	AngularAccelerationZ   float32        `json:"angularAccelerationZ"`   // Angular velocity z-component
	FrontWheelsAngle       float32        `json:"frontWheelsAngle"`       // Current front wheels angle in radians
}

func NewPacketCarMotionData(b []byte) *PacketCarMotionData {
	bc := utils.NewByteCursor(b)

	header := header.DecodeHeader(bc.B(24))
	CarMotionData := [22]CarMotionData{}

	for i := 0; i < 22; i++ {
		CarMotionData[i] = newCarMotionData(bc)
	}

	return &PacketCarMotionData{
		Header:                 header,
		CarMotionData:          CarMotionData,
		SuspensionPosition:     wheels.NewWheelsF(bc.B(4 * 4)),
		SuspensionVelocity:     wheels.NewWheelsF(bc.B(4 * 4)),
		SuspensionAcceleration: wheels.NewWheelsF(bc.B(4 * 4)),
		WheelSpeed:             wheels.NewWheelsF(bc.B(4 * 4)),
		WheelSlip:              wheels.NewWheelsF(bc.B(4 * 4)),
		LocalVelocityX:         bc.Float(),
		LocalVelocityY:         bc.Float(),
		LocalVelocityZ:         bc.Float(),
		AngularVelocityX:       bc.Float(),
		AngularVelocityY:       bc.Float(),
		AngularVelocityZ:       bc.Float(),
		AngularAccelerationX:   bc.Float(),
		AngularAccelerationY:   bc.Float(),
		AngularAccelerationZ:   bc.Float(),
		FrontWheelsAngle:       bc.Float(),
	}
}

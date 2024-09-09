package packets

import (
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/enum/gear"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/enum/mfd"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/types/header"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/types/wheels"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/utils"
)

type CarTelemetryData struct {
	Speed                   uint16          `json:"speed"`                   // Speed of car in kilometres per hour
	Throttle                float32         `json:"throttle"`                // Amount of throttle applied (0.0 to 1.0)
	Steer                   float32         `json:"steer"`                   // Steering (-1.0 (full lock left) to 1.0 (full lock right))
	Brake                   float32         `json:"brake"`                   // Amount of brake applied (0.0 to 1.0)
	Clutch                  uint8           `json:"clutch"`                  // Amount of clutch applied (0 to 100)
	Gear                    gear.Gear       `json:"gear"`                    // Gear selected (1-8, N=0, R=-1)
	EngineRPM               uint16          `json:"engineRPM"`               // Engine RPM
	Drs                     uint8           `json:"drs"`                     // 0 = off, 1 = on
	RevLightsPercent        uint8           `json:"revLightsPercent"`        // Rev lights indicator (percentage)
	RevLightsBitValue       uint16          `json:"revLightsBitValue"`       // Rev lights (bit 0 = leftmost LED, bit 14 = rightmost LED)
	BrakesTemperature       wheels.Wheels16 `json:"brakesTemperature"`       // Brakes temperature (celsius)
	TyresSurfaceTemperature wheels.Wheels8  `json:"tyresSurfaceTemperature"` // Tyres surface temperature (celsius)
	TyresInnerTemperature   wheels.Wheels8  `json:"tyresInnerTemperature"`   // Tyres inner temperature (celsius)
	EngineTemperature       uint16          `json:"engineTemperature"`       // Engine temperature (celsius)
	TyresPressure           wheels.WheelsF  `json:"tyresPressure"`           // Tyres pressure (PSI)
	SurfaceType             wheels.Wheels8  `json:"surfaceType"`             // Driving surface, see appendices
}

func newCarTelemetryData(bc *utils.ByteCursor) CarTelemetryData {
	return CarTelemetryData{
		Speed:                   bc.Uint16(),
		Throttle:                bc.Float(),
		Steer:                   bc.Float(),
		Brake:                   bc.Float(),
		Clutch:                  bc.Uint8(),
		Gear:                    gear.Gear(bc.Int8()),
		EngineRPM:               bc.Uint16(),
		Drs:                     bc.Uint8(),
		RevLightsPercent:        bc.Uint8(),
		RevLightsBitValue:       bc.Uint16(),
		BrakesTemperature:       wheels.NewWheels16(bc.B(2 * 4)),
		TyresSurfaceTemperature: wheels.NewWheels8(bc.B(4)),
		TyresInnerTemperature:   wheels.NewWheels8(bc.B(4)),
		EngineTemperature:       bc.Uint16(),
		TyresPressure:           wheels.NewWheelsF(bc.B(4 * 4)),
		SurfaceType:             wheels.NewWheels8(bc.B(4)),
	}
}

type PacketCarTelemetryData struct {
	Header header.Header `json:"header"` // Header

	CarTelemetryData [22]CarTelemetryData `json:"carTelemetryData"`

	MfdPanelIndex                mfd.Mfd `json:"mfdPanelIndex"`                // Index of MFD panel open - 255 = MFD closed Single player, race – 0 = Car setup, 1 = Pits 2 = Damage, 3 =  Engine, 4 = Temperatures May vary depending on game mode
	MfdPanelIndexSecondaryPlayer mfd.Mfd `json:"mfdPanelIndexSecondaryPlayer"` // See above
	SuggestedGear                int8    `json:"suggestedGear"`                // Suggested gear for the player (1-8) 0 if no gear suggested
}

func NewPacketCarTelemetryData(b []byte) *PacketCarTelemetryData {
	bc := utils.NewByteCursor(b)

	header := header.DecodeHeader(bc.B(24))
	carTelemetryData := [22]CarTelemetryData{}

	for i := 0; i < 22; i++ {
		carTelemetryData[i] = newCarTelemetryData(bc)
	}

	return &PacketCarTelemetryData{
		Header:                       header,
		CarTelemetryData:             carTelemetryData,
		MfdPanelIndex:                mfd.Mfd(bc.Uint8()),
		MfdPanelIndexSecondaryPlayer: mfd.Mfd(bc.Uint8()),
		SuggestedGear:                bc.Int8(),
	}
}

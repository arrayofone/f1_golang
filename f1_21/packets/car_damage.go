package packets

import (
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/types/header"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/types/wheels"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/utils"
)

type CarDamageData struct {
	TyresWear            wheels.WheelsF `json:"tyresWear"`            // Tyre wear (percentage)
	TyresDamage          wheels.Wheels8 `json:"tyresDamage"`          // Tyre damage (percentage)
	BrakesDamage         wheels.Wheels8 `json:"brakesDamage"`         // Brakes damage (percentage)
	FrontLeftWingDamage  uint8          `json:"frontLeftWingDamage"`  // Front left wing damage (percentage)
	FrontRightWingDamage uint8          `json:"frontRightWingDamage"` // Front right wing damage (percentage)
	RearWingDamage       uint8          `json:"rearWingDamage"`       // Rear wing damage (percentage)
	FloorDamage          uint8          `json:"floorDamage"`          // Floor damage (percentage)
	DiffuserDamage       uint8          `json:"diffuserDamage"`       // Diffuser damage (percentage)
	SidepodDamage        uint8          `json:"sidepodDamage"`        // Sidepod damage (percentage)
	DrsFault             uint8          `json:"drsFault"`             // Indicator for DRS fault, 0 = OK, 1 = fault
	GearBoxDamage        uint8          `json:"gearBoxDamage"`        // Gear box damage (percentage)
	EngineDamage         uint8          `json:"engineDamage"`         // Engine damage (percentage)
	EngineMGUHWear       uint8          `json:"engineMGUHWear"`       // Engine wear MGU-H (percentage)
	EngineESWear         uint8          `json:"engineESWear"`         // Engine wear ES (percentage)
	EngineCEWear         uint8          `json:"engineCEWear"`         // Engine wear CE (percentage)
	EngineICEWear        uint8          `json:"engineICEWear"`        // Engine wear ICE (percentage)
	EngineMGUKWear       uint8          `json:"engineMGUKWear"`       // Engine wear MGU-K (percentage)
	EngineTCWear         uint8          `json:"engineTCWear"`         // Engine wear TC (percentage)
}

func newCarDamageData(bc *utils.ByteCursor) CarDamageData {
	return CarDamageData{
		TyresWear:            wheels.NewWheelsF(bc.B(4 * 4)),
		TyresDamage:          wheels.NewWheels8(bc.B(4)),
		BrakesDamage:         wheels.NewWheels8(bc.B(4)),
		FrontLeftWingDamage:  bc.Uint8(),
		FrontRightWingDamage: bc.Uint8(),
		RearWingDamage:       bc.Uint8(),
		FloorDamage:          bc.Uint8(),
		DiffuserDamage:       bc.Uint8(),
		SidepodDamage:        bc.Uint8(),
		DrsFault:             bc.Uint8(),
		GearBoxDamage:        bc.Uint8(),
		EngineDamage:         bc.Uint8(),
		EngineMGUHWear:       bc.Uint8(),
		EngineESWear:         bc.Uint8(),
		EngineCEWear:         bc.Uint8(),
		EngineICEWear:        bc.Uint8(),
		EngineMGUKWear:       bc.Uint8(),
		EngineTCWear:         bc.Uint8(),
	}
}

type PacketCarDamageData struct {
	Header header.Header `json:"header"` // Header

	CarDamageData [22]CarDamageData `json:"carDamageData"`
}

func NewPacketCarDamageData(b []byte) *PacketCarDamageData {
	bc := utils.NewByteCursor(b)

	header := header.DecodeHeader(bc.B(24))
	CarDamageData := [22]CarDamageData{}

	for i := 0; i < 22; i++ {
		CarDamageData[i] = newCarDamageData(bc)
	}

	return &PacketCarDamageData{
		Header:        header,
		CarDamageData: CarDamageData,
	}
}

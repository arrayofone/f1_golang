package wheels

import (
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/utils"
)

type WheelsF struct {
	RL float32 `json:"rl"`
	RR float32 `json:"rr"`
	FL float32 `json:"fl"`
	FR float32 `json:"fr"`
}

func NewWheelsF(b []byte) WheelsF {
	bc := utils.NewByteCursor(b)

	return WheelsF{
		RL: bc.Float(),
		RR: bc.Float(),
		FL: bc.Float(),
		FR: bc.Float(),
	}
}

type Wheels8 struct {
	RL uint8 `json:"rl"`
	RR uint8 `json:"rr"`
	FL uint8 `json:"fl"`
	FR uint8 `json:"fr"`
}

func NewWheels8(b []byte) Wheels8 {
	bc := utils.NewByteCursor(b)

	return Wheels8{
		RL: bc.Uint8(),
		RR: bc.Uint8(),
		FL: bc.Uint8(),
		FR: bc.Uint8(),
	}
}

type Wheels16 struct {
	RL uint16 `json:"rl"`
	RR uint16 `json:"rr"`
	FL uint16 `json:"fl"`
	FR uint16 `json:"fr"`
}

func NewWheels16(b []byte) Wheels16 {
	bc := utils.NewByteCursor(b)

	return Wheels16{
		RL: bc.Uint16(),
		RR: bc.Uint16(),
		FL: bc.Uint16(),
		FR: bc.Uint16(),
	}
}

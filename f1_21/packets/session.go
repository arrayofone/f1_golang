package packets

import (
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/enum/flag"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/enum/formula"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/enum/sessionType"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/enum/track"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/enum/weather"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/types/header"
	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21/utils"
)

type MarshalZone struct { //5 bytes
	ZoneStart float32   `json:"zoneStart"` // Fraction (0..1) of way through the lap the marshal zone starts
	ZoneFlag  flag.Flag `json:"zoneFlag"`  // -1 = invalid/unknown, 0 = none, 1 = green, 2 = blue, 3 = yellow, 4 = red
}

func newMarshalZone(bc *utils.ByteCursor) MarshalZone {
	return MarshalZone{
		ZoneStart: bc.Float(),
		ZoneFlag:  flag.Flag(bc.Int8()),
	}
}

type WeatherForecastSample struct {
	SessionType            sessionType.SessionType `json:"sessionType"`            // 0 = unknown, 1 = P1, 2 = P2, 3 = P3, 4 = Short P, 5 = Q1 6 = Q2, 7 = Q3, 8 = Short Q, 9 = OSQ, 10 = R, 11 = R2 12 = Time Trial
	TimeOffset             uint8                   `json:"timeOffset"`             // Time in minutes the forecast is for
	Weather                weather.Weather         `json:"weather"`                // Weather - 0 = clear, 1 = light cloud, 2 = overcast 3 = light rain, 4 = heavy rain, 5 = storm
	TrackTemperature       int8                    `json:"trackTemperature"`       // Track temp. in degrees Celsius
	TrackTemperatureChange int8                    `json:"trackTemperatureChange"` // Track temp. change – 0 = up, 1 = down, 2 = no change
	AirTemperature         int8                    `json:"airTemperature"`         // Air temp. in degrees celsius
	AirTemperatureChange   int8                    `json:"airTemperatureChange"`   // Air temp. change – 0 = up, 1 = down, 2 = no change
	RainPercentage         uint8                   `json:"rainPercentage"`         // Rain percentage (0-100)
}

func newWeatherForecastSample(bc *utils.ByteCursor) WeatherForecastSample {
	return WeatherForecastSample{
		SessionType:            sessionType.SessionType(bc.Uint8()),
		TimeOffset:             bc.Uint8(),
		Weather:                weather.Weather(bc.Uint8()),
		TrackTemperature:       bc.Int8(),
		TrackTemperatureChange: bc.Int8(),
		AirTemperature:         bc.Int8(),
		AirTemperatureChange:   bc.Int8(),
		RainPercentage:         bc.Uint8(),
	}
}

type PacketSessionData struct {
	Header header.Header `json:"header"` // Header

	Weather                   weather.Weather         `json:"weather"`                   // Weather - 0 = clear, 1 = light cloud, 2 = overcast  3 = light rain, 4 = heavy rain, 5 = storm
	TrackTemperature          int8                    `json:"trackTemperature"`          // Track temp. in degrees celsius
	AirTemperature            int8                    `json:"airTemperature"`            // Air temp. in degrees celsius
	TotalLaps                 uint8                   `json:"totalLaps"`                 // Total number of laps in this race
	TrackLength               uint16                  `json:"trackLength"`               // Track length in metres
	SessionType               sessionType.SessionType `json:"sessionType"`               // 0 = unknown, 1 = P1, 2 = P2, 3 = P3, 4 = Short P  5 = Q1, 6 = Q2, 7 = Q3, 8 = Short Q, 9 = OSQ  10 = R, 11 = R2, 12 = R3, 13 = Time Trial
	TrackId                   track.Track             `json:"trackId"`                   // -1 for unknown, 0-21 for tracks, see appendix
	Formula                   formula.Formula         `json:"formula"`                   // Formula, 0 = F1 Modern, 1 = F1 Classic, 2 = F2, 3 = F1 Generic
	SessionTimeLeft           uint16                  `json:"sessionTimeLeft"`           // Time left in session in seconds
	SessionDuration           uint16                  `json:"sessionDuration"`           // Session duration in seconds
	PitSpeedLimit             uint8                   `json:"pitSpeedLimit"`             // Pit speed limit in kilometres per hour
	GamePaused                bool                    `json:"gamePaused"`                // Whether the game is paused
	IsSpectating              bool                    `json:"isSpectating"`              // Whether the player is spectating
	SpectatorCarIndex         uint8                   `json:"spectatorCarIndex"`         // Index of the car being spectated
	SliProNativeSupport       uint8                   `json:"sliProNativeSupport"`       // SLI Pro support, 0 = inactive, 1 = active
	NumMarshalZones           uint8                   `json:"numMarshalZones"`           // Number of marshal zones to follow
	MarshalZones              []MarshalZone           `json:"narshalZones"`              // List of marshal zones – max 21
	SafetyCarStatus           uint8                   `json:"safetyCarStatus"`           // 0 = no safety car, 1 = full 2 = virtual, 3 = formation lap
	NetworkGame               bool                    `json:"networkGame"`               // 0 = offline, 1 = online
	NumWeatherForecastSamples uint8                   `json:"numWeatherForecastSamples"` // Number of weather samples to follow
	WeatherForecastSamples    []WeatherForecastSample `json:"weatherForecastSamples"`    // Array of weather forecast samples
	ForecastAccuracy          uint8                   `json:"forecastAccuracy"`          // 0 = Perfect, 1 = Approximate
	AiDifficulty              uint8                   `json:"aiDifficulty"`              // AI Difficulty rating – 0-110
	SeasonLinkIdentifier      uint32                  `json:"seasonLinkIdentifier"`      // Identifier for season - persists across saves
	WeekendLinkIdentifier     uint32                  `json:"weekendLinkIdentifier"`     // Identifier for weekend - persists across saves
	SessionLinkIdentifier     uint32                  `json:"sessionLinkIdentifier"`     // Identifier for session - persists across saves
	PitStopWindowIdealLap     uint8                   `json:"pitStopWindowIdealLap"`     // Ideal lap to pit on for current strategy (player)
	PitStopWindowLatestLap    uint8                   `json:"pitStopWindowLatestLap"`    // Latest lap to pit on for current strategy (player)
	PitStopRejoinPosition     uint8                   `json:"pitStopRejoinPosition"`     // Predicted position to rejoin at (player)
	SteeringAssist            bool                    `json:"steeringAssist"`            // 0 = off, 1 = on
	BrakingAssist             uint8                   `json:"brakingAssist"`             // 0 = off, 1 = low, 2 = medium, 3 = high
	GearboxAssist             uint8                   `json:"gearboxAssist"`             // 1 = manual, 2 = manual & suggested gear, 3 = auto
	PitAssist                 bool                    `json:"pitAssist"`                 // 0 = off, 1 = on
	PitReleaseAssist          bool                    `json:"pitReleaseAssist"`          // 0 = off, 1 = on
	ERSAssist                 bool                    `json:"ersAssist"`                 // 0 = off, 1 = on
	DRSAssist                 bool                    `json:"drsAssist"`                 // 0 = off, 1 = on
	DynamicRacingLine         uint8                   `json:"dynamicRacingLine"`         // 0 = off, 1 = corners only, 2 = full
	DynamicRacingLineType     uint8                   `json:"dynamicRacingLineType"`     // 0 = 2D, 1 = 3D
}

func NewPacketSessionData(b []byte) *PacketSessionData {
	bc := utils.NewByteCursor(b)

	header := header.DecodeHeader(bc.B(24))

	weather := weather.Weather(bc.Uint8())
	trackTemperature := bc.Int8()
	airTemperature := bc.Int8()
	totalLaps := bc.Uint8()
	trackLength := bc.Uint16()
	sessionType := sessionType.SessionType(bc.Uint8())
	trackId := track.Track(bc.Uint8())
	formula := formula.Formula(bc.Uint8())
	sessionTimeLeft := bc.Uint16()
	sessionDuration := bc.Uint16()
	pitSpeedLimit := bc.Uint8()
	gamePaused := bc.Bool()
	isSpectating := bc.Bool()
	spectatorCarIndex := bc.Uint8()
	sliProNativeSupport := bc.Uint8()
	numMarshalZones := bc.Uint8()
	marshalZones := make([]MarshalZone, numMarshalZones)

	for i := 0; uint8(i) < numMarshalZones; i++ {
		marshalZones[i] = newMarshalZone(bc)
	}

	safetyCarStatus := bc.Uint8()
	networkGame := bc.Bool()
	numWeatherForecastSamples := bc.Uint8()
	weatherForecastSamples := make([]WeatherForecastSample, numWeatherForecastSamples)

	for i := 0; uint8(i) < numWeatherForecastSamples; i++ {
		weatherForecastSamples[i] = newWeatherForecastSample(bc)
	}

	return &PacketSessionData{
		Header:                    header,
		Weather:                   weather,
		TrackTemperature:          trackTemperature,
		AirTemperature:            airTemperature,
		TotalLaps:                 totalLaps,
		TrackLength:               trackLength,
		SessionType:               sessionType,
		TrackId:                   trackId,
		Formula:                   formula,
		SessionTimeLeft:           sessionTimeLeft,
		SessionDuration:           sessionDuration,
		PitSpeedLimit:             pitSpeedLimit,
		GamePaused:                gamePaused,
		IsSpectating:              isSpectating,
		SpectatorCarIndex:         spectatorCarIndex,
		SliProNativeSupport:       sliProNativeSupport,
		NumMarshalZones:           numMarshalZones,
		MarshalZones:              marshalZones, //5 bytes * 21 objects
		SafetyCarStatus:           safetyCarStatus,
		NetworkGame:               networkGame,
		NumWeatherForecastSamples: numWeatherForecastSamples,
		WeatherForecastSamples:    weatherForecastSamples,
		ForecastAccuracy:          bc.Uint8(),
		AiDifficulty:              bc.Uint8(),
		SeasonLinkIdentifier:      bc.Uint32(),
		WeekendLinkIdentifier:     bc.Uint32(),
		SessionLinkIdentifier:     bc.Uint32(),
		PitStopWindowIdealLap:     bc.Uint8(),
		PitStopWindowLatestLap:    bc.Uint8(),
		PitStopRejoinPosition:     bc.Uint8(),
		SteeringAssist:            bc.Bool(),
		BrakingAssist:             bc.Uint8(),
		GearboxAssist:             bc.Uint8(),
		PitAssist:                 bc.Bool(),
		PitReleaseAssist:          bc.Bool(),
		ERSAssist:                 bc.Bool(),
		DRSAssist:                 bc.Bool(),
		DynamicRacingLine:         bc.Uint8(),
		DynamicRacingLineType:     bc.Uint8(),
	}
}

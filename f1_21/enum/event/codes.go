package event

type EventCode string

const (
	SessionStarted            EventCode = "SSTA" // Sent when the session starts
	SessionEnded              EventCode = "SEND" // Sent when the session ends
	FastestLap                EventCode = "FTLP" // When a driver achieves the fastest lap
	Retirement                EventCode = "RTMT" // When a driver retires
	DRSEnabled                EventCode = "DRSE" // Race control have enabled DRS
	DRSDisabled               EventCode = "DRSD" // Race control have disabled DRS
	TeammateInPits            EventCode = "TMPT" // Your team mate has entered the pits
	ChequeredFlag             EventCode = "CHQF" // The chequered flag has been waved
	RaceWinner                EventCode = "RCWN" // The race winner is announced
	PenaltyIssued             EventCode = "PENA" // A penalty has been issued – details in event
	SpeedTrapTriggered        EventCode = "SPTP" // Speed trap has been triggered by fastest speed
	StartLights               EventCode = "STLG" // Start lights – number shown
	LightsOut                 EventCode = "LGOT" // Lights out
	DriveThroughPenaltyServed EventCode = "DTSV" // Drive through penalty served
	StopGoPenaltyServed       EventCode = "SGSV" // Stop go penalty served
	Flashback                 EventCode = "FLBK" // Flashback activated
	ButtonStatus              EventCode = "BUTN" // Button status changed
)

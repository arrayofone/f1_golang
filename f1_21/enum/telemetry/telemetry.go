package telemetry

type TelemetryStatus uint8

const (
	Restricted TelemetryStatus = iota
	Public
)

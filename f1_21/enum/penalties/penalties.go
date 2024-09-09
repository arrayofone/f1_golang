package penalties

type Penalty uint8

const (
	Drive_through Penalty = iota
	Stop_Go
	Grid_penalty
	Penalty_reminder
	Time_penalty
	Warning
	Disqualified
	Removed_from_formation_lap
	Parked_too_long_timer
	Tyre_regulations
	This_lap_invalidated
	This_and_next_lap_invalidated
	This_lap_invalidated_without_reason
	This_and_next_lap_invalidated_without_reason
	This_and_previous_lap_invalidated
	This_and_previous_lap_invalidated_without_reason
	Retired
	Black_flag_timer
)

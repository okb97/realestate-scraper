package model

import "time"

type UsedCondosBusStop struct {
	ID               string    `db:"id"`
	UsedCondoID      int       `db:"used_condo_id"`
	BusStopID        string    `db:"bus_stop_id"`
	WalkingMinutes   int       `db:"walking_minutes"`
	DeleteFlag       bool      `db:"delete_flag"`
	RegisterDateTime time.Time `db:"register_datetime"`
	UpdateDateTime   time.Time `db:"update_datetime"`
	RegisterFunction string    `db:"register_function"`
	UpdateFunction   string    `db:"update_function"`
	ProcessID        *string   `db:"process_id"`
}

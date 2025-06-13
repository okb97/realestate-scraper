package model

import "time"

type UsedCondosBusStop struct {
	ID               int       `db:"id"`
	UsedCondoID      int       `db:"used_condo_id"`
	BusStopName      string    `db:"bus_stop_name"`
	StationName      string    `db:"station_name"`
	TrainLineName    string    `db:"train_line_name"`
	BusMinutes       int       `db:"bus_minutes"`
	WalkingMinutes   int       `db:"walking_minutes"`
	DeleteFlag       bool      `db:"delete_flag"`
	RegisterDateTime time.Time `db:"register_date_time"`
	UpdateDateTime   time.Time `db:"update_date_time"`
	RegisterFunction string    `db:"register_function"`
	UpdateFunction   string    `db:"update_function"`
	ProcessID        *string   `db:"process_id"`
}

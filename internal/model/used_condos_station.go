package model

import "time"

type UsedCondosStation struct {
	ID               string    `db:"id"`
	UsedCondoID      int       `db:"used_condo_id"`
	StationID        string    `db:"station_id"`
	WalkingMinutes   int       `db:"walking_minutes"`
	DeleteFlag       bool      `db:"delete_flag"`
	RegisterDateTime time.Time `db:"register_datetime"`
	UpdateDateTime   time.Time `db:"update_datetime"`
	RegisterFunction string    `db:"register_function"`
	UpdateFunction   string    `db:"update_function"`
	ProcessID        string    `db:"process_id"`
}

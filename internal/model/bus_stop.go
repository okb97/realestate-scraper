package model

import "time"

type BusStop struct {
	BusStopID        string    `db:"bus_stop_id"`
	BusStopName      string    `db:"bus_stop_name"`
	BusLineID        string    `db:"bus_line_id"`
	AddressID        int       `db:"address_id"`
	DeleteFlag       bool      `db:"delete_flag"`
	RegisterDateTime time.Time `db:"register_datetime"`
	UpdateDateTime   time.Time `db:"update_datetime"`
	RegisterFunction string    `db:"register_function"`
	UpdateFunction   string    `db:"update_function"`
	ProcessID        *string   `db:"process_id"`
}

package model

import "time"

type Station struct {
	StationID        string    `db:"station_id"`
	TrainLineID      string    `db:"train_line_id"`
	StationName      string    `db:"station_name"`
	PostalCode       string    `db:"postal_code`
	AddressID        int       `db:"address_id"`
	Address          string    `db:"address"`
	DeleteFlag       bool      `db:"delete_flag"`
	RegisterDateTime time.Time `db:"register_datetime"`
	UpdateDateTime   time.Time `db:"update_datetime"`
	RegisterFunction string    `db:"register_function"`
	UpdateFunction   string    `db:"update_function"`
	ProcessID        *string   `db:"process_id"`
}

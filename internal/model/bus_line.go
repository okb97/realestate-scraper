package model

import "time"

type BusLine struct {
	BusLineID        string    `db:"bus_line_id"`
	BusCompanyName   string    `db:"bus_company_name"`
	DeleteFlag       bool      `db:"delete_flag"`
	RegisterDateTime time.Time `db:"register_datetime"`
	UpdateDateTime   time.Time `db:"update_datetime"`
	RegisterFunction string    `db:"register_function"`
	UpdateFunction   string    `db:"update_function"`
	ProcessID        *string   `db:"process_id"`
}

package model

import "time"

type Address struct {
	AddressID        string    `db:"address_id"`
	PostalCode       string    `db:"postalcode"`
	Prefecture       string    `db:"prefecture"`
	City             string    `db:"city"`
	Town             string    `db:"town"`
	DeleteFlag       bool      `db:"delete_flag"`
	RegisterDateTime time.Time `db:"register_datetime"`
	UpdateDateTime   time.Time `db:"update_datetime"`
	RegisterFunction string    `db:"register_function"`
	UpdateFunction   string    `db:"update_function"`
	ProcessID        *string   `db:"process_id"`
}

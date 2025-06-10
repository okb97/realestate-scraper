package model

import "time"

type UsedCondo struct {
	UsedCondoID             int       `db:"used_condo_id"`
	Url                     string    `db:"url"`
	UsedCondoName           string    `db:"UsedCondoName"`
	PriceText               string    `db:"price_text"`
	PriceNum                int       `db:"price_num"`
	Layout                  string    `db:"layout"`
	TotalUnitsText          string    `db:"total_units_text"`
	TotalUnitsNum           int       `db:"total_units_num"`
	OccupiedAreaText        string    `db:"occupied_area_text"`
	OccupiedAreaNum         float64   `db:"occupied_area_num"`
	OtherAreaText           string    `db:"other_area_text"`
	OtherAreaNum            float64   `db:"other_area_num"`
	BuiltAt                 time.Time `db:"built_at"`
	AddressID               int       `db:"address_id"`
	Address                 string    `db:"address"`
	SalesSchedule           string    `db:"sales_schedule"`
	EventInfo               string    `db:"event_info"`
	MostPriceRange          string    `db:"most_price_range"`
	MaintenanceFee          string    `db:"maintenance_fee"`
	RepairReserveFund       string    `db:"repair_reserve_fund"`
	IntialRepairReserveFund string    `db:"intial_repair_reserve_fund"`
	AdditionalCosts         string    `db:"additional_costs"`
	Floor                   string    `db:"floor"`
	Direction               string    `db:"direction"`
	EnergyEfficiency        string    `db:"energy_efficiency"`
	InsulationPerformance   string    `db:"insulation_performance"`
	EstimatedUtilityCost    string    `db:"estimated_utility_cost"`
	Reform                  string    `db:"reform"`
	Structure               string    `db:"structure"`
	SiteArea                string    `db:"SiteArea"`
	LandRight               string    `db:"land_right"`
	Zoning                  string    `db:"zoning"`
	Parking                 string    `db:"parking"`
	Contractor              string    `db:"contractor"`
	ListedFlag              bool      `db:"listed_flag"`
	DeleteFlag              bool      `db:"delete_flag"`
	RegisterDateTime        time.Time `db:"register_datetime"`
	UpdateDateTime          time.Time `db:"update_datetime"`
	RegisterFunction        string    `db:"register_function"`
	UpdateFunction          string    `db:"update_function"`
	ProcessID               *string   `db:"process_id"`
}

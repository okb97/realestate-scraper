package model

type DetailCollector struct {
	Url                     string   `db:"Url"`
	UsedCondoName           string   `db:"UsedCondoName"`
	PriceText               string   `db:"PriceText"`
	Layout                  string   `db:"Layout"`
	TotalUnitsText          string   `db:"TotalUnitsText"`
	OccupiedAreaText        string   `db:"OccupiedAreaText"`
	OtherAreaText           string   `db:"OtherAreaText"`
	BuiltAt                 string   `db:"BuiltAt"`
	Address                 string   `db:"Address"`
	SalesSchedule           string   `db:"SalesSchedule"`
	EventInfo               string   `db:"EventInfo"`
	MostPriceRange          string   `db:"MostPriceRange"`
	MaintenanceFee          string   `db:"MaintenanceFee"`
	RepairReserveFund       string   `db:"RepairReserveFund"`
	IntialRepairReserveFund string   `db:"IntialRepairReserveFund"`
	AdditionalCosts         string   `db:"AdditionalCosts"`
	Floor                   string   `db:"Floor"`
	Direction               string   `db:"Direction"`
	EnergyEfficiency        string   `db:"EnergyEfficiency"`
	InsulationPerformance   string   `db:"InsulationPerformance"`
	EstimatedUtilityCost    string   `db:"EstimatedUtilityCost"`
	Reform                  string   `db:"Reform"`
	Structure               string   `db:"Structure"`
	SiteArea                string   `db:"SiteArea"`
	LandRight               string   `db:"LandRight"`
	Zoning                  string   `db:"Zoning"`
	Parking                 string   `db:"Parking"`
	Contractor              string   `db:"Contractor"`
	Transportation          []string `db:"Transportation"`
	BusSlice                []string `db:"BusSlice"`
}

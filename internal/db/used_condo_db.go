package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"time"

	"github.com/okb97/realestate-scraper/internal/model"
)

var Conn *sql.DB

func InsertUsedCondo(ctx context.Context, tx *sql.Tx, condo model.UsedCondo) (int, bool, error) {
	query := `
WITH upsert AS (
    INSERT INTO used_condos (
        url,
			used_condo_name,
			price_text,
			price_num,
			layout,
			total_units_text,
			total_units_num,
			occupied_area_text,
			occupied_area_num,
			other_area_text,
			other_area_num,
			floor,
			built_at,
			address_id,
			address,
			sales_schedule,
			event_info,
			most_price_range,
			maintenance_fee,
			repair_reserve_fund,
			intial_repair_reserve_fund,
			additional_costs,
			direction,
			energy_efficiency,
			insulation_performance,
			estimated_utility_cost,
			reform,
			structure,
			site_area,
			land_right,
			zoning,
			parking,
			contractor,
			listed_flag,
			delete_flag,
			register_date_time,
			update_date_time,
			register_function,
			update_function
    )
    VALUES (
        $1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
			$11, $12, $13, $14, $15, $16, $17, $18,
			$19, $20, $21, $22, $23, $24, $25, $26,
			$27, $28, $29, $30, $31, $32, $33, $34,
			$35, $36, $37, $38, $39
	)
    ON CONFLICT (url) DO UPDATE SET
        used_condo_name = EXCLUDED.used_condo_name,
		price_text = EXCLUDED.price_text,
		price_num = EXCLUDED.price_num,
		layout = EXCLUDED.layout,
		total_units_text = EXCLUDED.total_units_text,
		total_units_num = EXCLUDED.total_units_num,
		occupied_area_text = EXCLUDED.occupied_area_text,
		occupied_area_num = EXCLUDED.occupied_area_num,
		other_area_text = EXCLUDED.other_area_text,
		other_area_num = EXCLUDED.other_area_num,
		floor = EXCLUDED.floor,
		built_at = EXCLUDED.built_at,
		address_id = EXCLUDED.address_id,
		address = EXCLUDED.address,
		sales_schedule = EXCLUDED.sales_schedule,
		event_info = EXCLUDED.event_info,
		most_price_range = EXCLUDED.most_price_range,
		maintenance_fee = EXCLUDED.maintenance_fee,
		repair_reserve_fund = EXCLUDED.repair_reserve_fund,
		intial_repair_reserve_fund = EXCLUDED.intial_repair_reserve_fund,
		additional_costs = EXCLUDED.additional_costs,
		direction = EXCLUDED.direction,
		energy_efficiency = EXCLUDED.energy_efficiency,
		insulation_performance = EXCLUDED.insulation_performance,
		estimated_utility_cost = EXCLUDED.estimated_utility_cost,
		reform = EXCLUDED.reform,
		structure = EXCLUDED.structure,
		site_area = EXCLUDED.site_area,
		land_right = EXCLUDED.land_right,
		zoning = EXCLUDED.zoning,
		parking = EXCLUDED.parking,
		contractor = EXCLUDED.contractor,
		listed_flag = EXCLUDED.listed_flag,
		update_date_time = NOW()
	WHERE
		used_condos.used_condo_name IS DISTINCT FROM EXCLUDED.used_condo_name OR
		used_condos.price_text IS DISTINCT FROM EXCLUDED.price_text OR
		used_condos.price_num IS DISTINCT FROM EXCLUDED.price_num OR
		used_condos.layout IS DISTINCT FROM EXCLUDED.layout OR
		used_condos.total_units_text IS DISTINCT FROM EXCLUDED.total_units_text OR
		used_condos.total_units_num IS DISTINCT FROM EXCLUDED.total_units_num OR
		used_condos.occupied_area_text IS DISTINCT FROM EXCLUDED.occupied_area_text OR
		used_condos.occupied_area_num IS DISTINCT FROM EXCLUDED.occupied_area_num OR
		used_condos.other_area_text IS DISTINCT FROM EXCLUDED.other_area_text OR
		used_condos.other_area_num IS DISTINCT FROM EXCLUDED.other_area_num OR
		used_condos.floor IS DISTINCT FROM EXCLUDED.floor OR
		used_condos.built_at IS DISTINCT FROM EXCLUDED.built_at OR
		used_condos.address_id IS DISTINCT FROM EXCLUDED.address_id OR
		used_condos.address IS DISTINCT FROM EXCLUDED.address OR
		used_condos.sales_schedule IS DISTINCT FROM EXCLUDED.sales_schedule OR
		used_condos.event_info IS DISTINCT FROM EXCLUDED.event_info OR
		used_condos.most_price_range IS DISTINCT FROM EXCLUDED.most_price_range OR
		used_condos.maintenance_fee IS DISTINCT FROM EXCLUDED.maintenance_fee OR
		used_condos.repair_reserve_fund IS DISTINCT FROM EXCLUDED.repair_reserve_fund OR
		used_condos.intial_repair_reserve_fund IS DISTINCT FROM EXCLUDED.intial_repair_reserve_fund OR
		used_condos.additional_costs IS DISTINCT FROM EXCLUDED.additional_costs OR
		used_condos.direction IS DISTINCT FROM EXCLUDED.direction OR
		used_condos.energy_efficiency IS DISTINCT FROM EXCLUDED.energy_efficiency OR
		used_condos.insulation_performance IS DISTINCT FROM EXCLUDED.insulation_performance OR
		used_condos.estimated_utility_cost IS DISTINCT FROM EXCLUDED.estimated_utility_cost OR
		used_condos.reform IS DISTINCT FROM EXCLUDED.reform OR
		used_condos.structure IS DISTINCT FROM EXCLUDED.structure OR
		used_condos.site_area IS DISTINCT FROM EXCLUDED.site_area OR
		used_condos.land_right IS DISTINCT FROM EXCLUDED.land_right OR
		used_condos.zoning IS DISTINCT FROM EXCLUDED.zoning OR
		used_condos.parking IS DISTINCT FROM EXCLUDED.parking OR
		used_condos.contractor IS DISTINCT FROM EXCLUDED.contractor OR
		used_condos.listed_flag IS DISTINCT FROM EXCLUDED.listed_flag
    RETURNING used_condo_id
)
SELECT used_condo_id FROM upsert
LIMIT 1
    `
	var insertedID int

	err := tx.QueryRowContext(ctx, query,
		condo.Url,
		condo.UsedCondoName,
		condo.PriceText,
		int(condo.PriceNum),
		condo.Layout,
		condo.TotalUnitsText,
		condo.TotalUnitsNum,
		condo.OccupiedAreaText,
		condo.OccupiedAreaNum,
		condo.OtherAreaText,
		condo.OtherAreaNum,
		condo.Floor,
		condo.BuiltAt,
		condo.AddressID,
		condo.Address,
		condo.SalesSchedule,
		condo.EventInfo,
		condo.MostPriceRange,
		condo.MaintenanceFee,
		condo.RepairReserveFund,
		condo.IntialRepairReserveFund,
		condo.AdditionalCosts,
		condo.Direction,
		condo.EnergyEfficiency,
		condo.InsulationPerformance,
		condo.EstimatedUtilityCost,
		condo.Reform,
		condo.Structure,
		condo.SiteArea,
		condo.LandRight,
		condo.Zoning,
		condo.Parking,
		condo.Contractor,
		condo.ListedFlag,
		false, // delete_flag
		time.Now(),
		time.Now(),
		"used_condo",
		"used_condo",
	).Scan(&insertedID)

	if err != nil {
		if err == sql.ErrNoRows {
			// 既存で更新もされていない場合
			return 0, false, nil
		}
		log.Printf("InsertUsedCondo: URL検索失敗: %v", err)
		return 0, false, err
	}

	// 挿入または更新された場合
	return insertedID, true, nil

}

type UsedCondoWithDetail struct {
	UsedCondo model.UsedCondo
	Address   struct {
		PostalCode string
		Prefecture string
		City       string
		Town       string
	}
	Stations []struct {
		StationID      int    `json:"station_id"`
		StationName    string `json:"station_name"`
		WalkingMinutes int    `json:"walking_minutes"`
	}
	BusStops []struct {
		BusStopName    string `json:"bus_stop_name"`
		StationName    string `json:"station_name"`
		TrainLineName  string `json:"train_line_name"`
		BusMinutes     int    `json:"bus_minutes"`
		WalkingMinutes int    `json:"walking_minutes"`
	}
}

func GetAllUsedCondos() ([]UsedCondoWithDetail, error) {
	query := `
	SELECT
	  uc.used_condo_id,
	  uc.url,
	  uc.used_condo_name,
	  uc.price_text,
	  uc.price_num,
	  uc.layout,
	  uc.total_units_text,
	  uc.total_units_num,
	  uc.occupied_area_text,
	  uc.occupied_area_num,
	  uc.other_area_text,
	  uc.other_area_num,
	  uc.floor,
	  uc.built_at,
	  uc.address_id,
	  a.postal_code,
	  a.prefecture,
	  a.city,
	  a.town,
	  uc.sales_schedule,
	  uc.event_info,
	  uc.most_price_range,
	  uc.maintenance_fee,
	  uc.repair_reserve_fund,
	  uc.intial_repair_reserve_fund,
	  uc.additional_costs,
	  uc.direction,
	  uc.energy_efficiency,
	  uc.insulation_performance,
	  uc.estimated_utility_cost,
	  uc.reform,
	  uc.structure,
	  uc.site_area,
	  uc.land_right,
	  uc.zoning,
	  uc.parking,
	  uc.contractor,
	  uc.listed_flag,
	  uc.delete_flag,
	  uc.register_date_time,
	  uc.update_date_time,
	  uc.register_function,
	  uc.update_function,
	  uc.process_id,
	  COALESCE(json_agg(
		json_build_object(
		  'station_id', s.station_id,
		  'station_name', s.station_name,
		  'walking_minutes', ucs.walking_minutes
		)
	  ) FILTER (WHERE s.station_id IS NOT NULL), '[]') AS stations,
	  COALESCE(json_agg(
		DISTINCT jsonb_build_object(
		  'bus_stop_name', ucb.bus_stop_name,
		  'station_name', ucb.station_name,
		  'train_line_name', ucb.train_line_name,
		  'bus_minutes', ucb.bus_minutes,
		  'walking_minutes', ucb.walking_minutes
		)
	  ) FILTER (WHERE ucb.bus_stop_name IS NOT NULL), '[]') AS bus_stops
	FROM
	  used_condos uc
	LEFT JOIN address a ON uc.address_id = a.address_id
	LEFT JOIN used_condos_stations ucs ON uc.used_condo_id = ucs.used_condo_id
	LEFT JOIN station s ON ucs.station_id = s.station_id
	LEFT JOIN used_condos_bus_stop ucb ON uc.used_condo_id = ucb.used_condo_id
	GROUP BY
	  uc.used_condo_id,
	  a.address_id
	ORDER BY
	  uc.used_condo_id;
	`
	log.Println("DB Query開始")
	rows, err := Conn.Query(query)
	if err != nil {
		log.Printf("DB Query error: %v", err)
		return nil, err
	}
	defer rows.Close()
	log.Println("DB Query成功")

	var condos []UsedCondoWithDetail
	for rows.Next() {
		var c UsedCondoWithDetail
		var stationsJSON, busStopsJSON []byte

		// UsedCondoのフィールドを1つずつ渡す
		err := rows.Scan(
			&c.UsedCondo.UsedCondoID,
			&c.UsedCondo.Url,
			&c.UsedCondo.UsedCondoName,
			&c.UsedCondo.PriceText,
			&c.UsedCondo.PriceNum,
			&c.UsedCondo.Layout,
			&c.UsedCondo.TotalUnitsText,
			&c.UsedCondo.TotalUnitsNum,
			&c.UsedCondo.OccupiedAreaText,
			&c.UsedCondo.OccupiedAreaNum,
			&c.UsedCondo.OtherAreaText,
			&c.UsedCondo.OtherAreaNum,
			&c.UsedCondo.Floor,
			&c.UsedCondo.BuiltAt,
			&c.UsedCondo.AddressID,
			&c.Address.PostalCode,
			&c.Address.Prefecture,
			&c.Address.City,
			&c.Address.Town,
			&c.UsedCondo.SalesSchedule,
			&c.UsedCondo.EventInfo,
			&c.UsedCondo.MostPriceRange,
			&c.UsedCondo.MaintenanceFee,
			&c.UsedCondo.RepairReserveFund,
			&c.UsedCondo.IntialRepairReserveFund,
			&c.UsedCondo.AdditionalCosts,
			&c.UsedCondo.Direction,
			&c.UsedCondo.EnergyEfficiency,
			&c.UsedCondo.InsulationPerformance,
			&c.UsedCondo.EstimatedUtilityCost,
			&c.UsedCondo.Reform,
			&c.UsedCondo.Structure,
			&c.UsedCondo.SiteArea,
			&c.UsedCondo.LandRight,
			&c.UsedCondo.Zoning,
			&c.UsedCondo.Parking,
			&c.UsedCondo.Contractor,
			&c.UsedCondo.ListedFlag,
			&c.UsedCondo.DeleteFlag,
			&c.UsedCondo.RegisterDateTime,
			&c.UsedCondo.UpdateDateTime,
			&c.UsedCondo.RegisterFunction,
			&c.UsedCondo.UpdateFunction,
			&c.UsedCondo.ProcessID,
			&stationsJSON,
			&busStopsJSON,
		)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(stationsJSON, &c.Stations); err != nil {
			log.Printf("stations JSON Unmarshal error: %v", err)
		}
		if err := json.Unmarshal(busStopsJSON, &c.BusStops); err != nil {
			log.Printf("busStops JSON Unmarshal error: %v", err)
		}

		condos = append(condos, c)
	}
	return condos, nil
}

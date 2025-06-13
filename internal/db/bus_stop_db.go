package db

import (
	"database/sql"
	"fmt"
	"time"
)

func InsertBusStopAccess(tx *sql.Tx, busStopName, stationName, trainLineName string, usedCondoID, busMin, walkMin int) error {
	now := time.Now()
	_, err := tx.Exec(`
	INSERT INTO used_condos_bus_stop (
		used_condo_id, bus_stop_name, station_name, train_line_name, bus_minutes, walking_minutes, 
		delete_flag, register_date_time, update_date_time, register_function, update_function, process_id
	) VALUES ($1, $2, $3, $4, $5, $6, false, $7, $7, 'bus_stop', 'bus_stop', 'scraper')
	`, usedCondoID, busStopName, stationName, trainLineName, busMin, walkMin, now)

	if err != nil {
		return fmt.Errorf("バス停情報挿入失敗: %w", err)
	}

	return nil
}

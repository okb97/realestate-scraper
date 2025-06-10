package db

import (
	"database/sql"
	"fmt"
	"time"
)

func InsertBusStopAccess(tx *sql.Tx, busStopName string, usedCondoID, walkMin int) error {
	var busStopID int
	err := tx.QueryRow(`
	SELECT bus_stop_id FROM bus_stop
	WHERE bus_stop_name ILIKE $1 LIMIT 1
	`, "%"+busStopName+"%").Scan(&busStopID)

	if err != nil {
		return fmt.Errorf("バス停検索失敗: %w", err)
	}

	now := time.Now()
	_, err = tx.Exec(`
	INSERT INTO used_condos_bus_stop (
		used_condo_id, bus_stop_id, walking_minutes, delete_flag,
		register_datetime, update_datetime, register_function, update_function
	) VALUES ($1, $2, $3, false, $4, $4, 'bus_stop', 'bus_stop')
	`, usedCondoID, busStopID, walkMin, now)
	return err
}

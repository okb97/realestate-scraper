package db

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"
)

func InsertStationAccess(tx *sql.Tx, trainLineName, stationName string, usedCondoID, walkMin int) error {
	var stationID string
	// 路線名の候補を取得（同義語を含む）
	lineNameCandidates := getLineNameCandidates(trainLineName)
	stationName = NormalizeNames(stationName)

	// プレースホルダとクエリ構築
	placeholders := ""
	args := []interface{}{"%" + stationName + "%"}
	for i, name := range lineNameCandidates {
		placeholders += fmt.Sprintf("t.train_line_name ILIKE $%d", i+2)
		args = append(args, "%"+name+"%")
		if i < len(lineNameCandidates)-1 {
			placeholders += " OR "
		}
	}

	query := fmt.Sprintf(`
	SELECT s.station_id FROM station s
	JOIN train_line t ON s.train_line_id = t.train_line_id
	WHERE s.station_name ILIKE $1 AND (%s)
	LIMIT 1
	`, placeholders)

	err := tx.QueryRow(query, args...).Scan(&stationID)
	if err != nil {
		log.Printf("駅検索失敗: %v (station: %s, line: %s)", err, stationName, trainLineName)
		return fmt.Errorf("駅検索失敗: %w", err)
	}

	now := time.Now()
	_, err = tx.Exec(`
	INSERT INTO used_condos_stations (
		used_condo_id, station_id, walking_minutes, delete_flag,
		register_date_time, update_date_time, register_function, update_function
	) VALUES ($1, $2, $3, false, $4, $4, 'station', 'station')
	`, usedCondoID, stationID, walkMin, now)
	if err != nil {
		log.Printf("INSERT失敗: %v (used_condo_id: %d, station_id: %v, walk_min: %d)", err, usedCondoID, stationID, walkMin)
	}
	return err
}

func getLineNameCandidates(name string) []string {
	switch name {
	case "JR京浜東北線":
		return []string{"JR京浜東北線", "JR根岸線"}
	case "JR根岸線":
		return []string{"JR根岸線", "JR京浜東北線"}
	default:
		return []string{name}
	}
}

func NormalizeNames(name string) string {
	replacer := strings.NewReplacer(
		"ケ", "ヶ",
	)

	normalizedName := replacer.Replace(name)

	return normalizedName
}

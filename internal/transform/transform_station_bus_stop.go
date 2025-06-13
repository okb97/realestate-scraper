package transform

import (
	"database/sql"
	"fmt"
	"regexp"
	"strings"

	"github.com/okb97/realestate-scraper/internal/db"
)

type AccessInfo struct {
	IsStation      bool
	TrainLineName  string // 路線名
	StationName    string // 駅名
	BusStopName    string // バス停名
	BusMinutes     int    // バス乗車時間
	WalkingMinutes int    // 徒歩時間
}

func TransformStationAndBusStop(tx *sql.Tx, transportations []string, usedCondoId int) error {
	for _, t := range transportations {
		accessList := parseAccessInfo(t)

		for _, access := range accessList {
			if access.IsStation {
				err := db.InsertStationAccess(tx, access.TrainLineName, access.StationName, usedCondoId, access.WalkingMinutes)
				if err != nil {
					return fmt.Errorf("駅の挿入失敗（%s %s）: %w", access.TrainLineName, access.StationName, err)
				}
			} else {
				err := db.InsertBusStopAccess(tx, access.BusStopName, access.StationName, access.TrainLineName, usedCondoId, access.BusMinutes, access.WalkingMinutes)
				if err != nil {
					return fmt.Errorf("バス停の挿入失敗（%s）: %w", access.BusStopName, err)
				}
			}
		}
	}
	return nil
}

func parseAccessInfo(transportation string) []AccessInfo {
	results := []AccessInfo{}

	// 複数の交通情報が含まれている場合の分割処理
	segments := splitTransportationString(transportation)

	for _, segment := range segments {
		segment = strings.TrimSpace(segment)
		if segment == "" {
			continue
		}

		// パターン1: 路線「駅」バス○分バス停名歩○分
		re1 := regexp.MustCompile(`(.+?)「(.+?)」バス(\d+)分(.+?)歩(\d+)分`)
		matches1 := re1.FindAllStringSubmatch(segment, -1)

		for _, m := range matches1 {
			trainLine := normalize(m[1])
			stationName := normalize(m[2])
			busMin := toInt(m[3])
			busStopName := normalize(m[4])
			walkMin := toInt(m[5])

			results = append(results, AccessInfo{
				IsStation:      false,
				TrainLineName:  trainLine,
				StationName:    stationName,
				BusStopName:    busStopName,
				BusMinutes:     busMin,
				WalkingMinutes: walkMin,
			})
		}
		if len(matches1) > 0 {
			continue
		}

		// パターン2: バス会社名「バス停名」歩○分
		re2 := regexp.MustCompile(`(.+バス)「(.+?)」(?:歩(\d+)分)?`)
		matches2 := re2.FindAllStringSubmatch(segment, -1)

		for _, m := range matches2 {
			busCompany := normalize(m[1])
			busStopName := normalize(m[2])
			walkMin := 0
			if len(m) > 3 && m[3] != "" {
				walkMin = toInt(m[3])
			}

			results = append(results, AccessInfo{
				IsStation:      false,
				TrainLineName:  busCompany,
				StationName:    "",
				BusStopName:    busStopName,
				BusMinutes:     0,
				WalkingMinutes: walkMin,
			})
		}
		if len(matches2) > 0 {
			continue
		}

		// パターン3: 通常の駅アクセス（路線「駅」歩○分）※路線名のみをマッチ
		re3 := regexp.MustCompile(`(.+線)「(.+?)」(?:歩(\d+)分)?`)
		matches3 := re3.FindAllStringSubmatch(segment, -1)

		for _, m := range matches3 {
			trainLine := normalize(m[1])
			stationName := normalize(m[2])
			walkMin := 0
			if len(m) > 3 && m[3] != "" {
				walkMin = toInt(m[3])
			}

			results = append(results, AccessInfo{
				IsStation:      true,
				TrainLineName:  trainLine,
				StationName:    stationName,
				BusStopName:    "",
				BusMinutes:     0,
				WalkingMinutes: walkMin,
			})
		}
	}

	return results
}

// 複雑な交通情報文字列をセグメントに分割する関数
func splitTransportationString(transportation string) []string {
	// スペースや改行で区切られた交通情報を分割
	segments := []string{}

	// 改行やスペースで分割
	parts := regexp.MustCompile(`[\s\n\r]+`).Split(transportation, -1)

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		// 複数の路線情報が連結されている場合を分割
		// 例: "東急田園都市線「南町田グランベリーパーク」歩22分 東急田園都市線「南町田グランベリーパーク」バス9分マークスプリングス歩1分"

		// 路線名「駅名」で始まる部分を検索して分割
		re := regexp.MustCompile(`([^\s]+線「[^」]+」[^\s]*)`)
		matches := re.FindAllString(part, -1)

		if len(matches) > 1 {
			// 複数のマッチがある場合は分割
			segments = append(segments, matches...)
		} else {
			// 単一のセグメントとして追加
			segments = append(segments, part)
		}
	}

	return segments
}

func normalize(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "　", "")
	s = strings.ReplaceAll(s, "ＪＲ", "JR")
	return s
}

func toInt(s string) int {
	var n int
	fmt.Sscanf(s, "%d", &n)
	return n
}

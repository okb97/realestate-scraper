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
	LineOrBusName  string
	PlaceName      string
	WalkingMinutes int
}

func TransformStationAndBusStop(tx *sql.Tx, transportations []string, usedCondoId int) error {
	for _, t := range transportations {
		accessList := parseAccessInfo(t)
		for _, access := range accessList {
			if access.IsStation {
				err := db.InsertStationAccess(tx, access.LineOrBusName, access.PlaceName, usedCondoId, access.WalkingMinutes)
				if err != nil {
					return fmt.Errorf("駅の挿入失敗（%s %s）: %w", access.LineOrBusName, access.PlaceName, err)
				}
			} else {
				err := db.InsertBusStopAccess(tx, access.PlaceName, usedCondoId, access.WalkingMinutes)
				if err != nil {
					return fmt.Errorf("バス停の挿入失敗（%s）: %w", access.PlaceName, err)
				}
			}
		}
	}
	return nil
}

func parseAccessInfo(transportation string) []AccessInfo {
	// 駅 or バス混在パターン
	re := regexp.MustCompile(`(?P<line>.+?)「(?P<name>.+?)」(?:バス(?P<busmin>\d+)分)?(?:歩(?P<walk>\d+)分)?`)
	matches := re.FindAllStringSubmatch(transportation, -1)
	results := []AccessInfo{}

	for _, m := range matches {
		line := normalize(m[1])
		name := normalize(m[2])
		busMin := m[3]
		walkMin := m[4]

		if busMin != "" {
			// バス停処理（バス停 + 歩）
			walk := toInt(walkMin)
			results = append(results, AccessInfo{
				IsStation:      false,
				LineOrBusName:  line,
				PlaceName:      name,
				WalkingMinutes: walk,
			})
		} else {
			// 駅処理（路線 + 駅 + 歩）
			walk := toInt(walkMin)
			results = append(results, AccessInfo{
				IsStation:      true,
				LineOrBusName:  line,
				PlaceName:      name,
				WalkingMinutes: walk,
			})
		}
	}
	return results
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

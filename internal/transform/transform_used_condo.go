package transform

import (
	"database/sql"
	"regexp"
	"strconv"
	"time"

	"github.com/okb97/realestate-scraper/internal/model"
	"github.com/okb97/realestate-scraper/internal/utils"
)

func TransformUsedCondo(tx *sql.Tx, detailCollectormodel model.DetailCollector) model.UsedCondo {
	var usedCondoModel model.UsedCondo

	usedCondoModel.Url = detailCollectormodel.Url
	usedCondoModel.UsedCondoName = detailCollectormodel.UsedCondoName
	usedCondoModel.PriceText = detailCollectormodel.PriceText
	usedCondoModel.PriceNum = priceTextToNum(detailCollectormodel.PriceText)
	usedCondoModel.Layout = detailCollectormodel.Layout
	usedCondoModel.TotalUnitsText = detailCollectormodel.TotalUnitsText
	usedCondoModel.TotalUnitsNum = totalUnitsTextToNum(detailCollectormodel.TotalUnitsText)
	usedCondoModel.OccupiedAreaText = detailCollectormodel.OccupiedAreaText
	usedCondoModel.OccupiedAreaNum = areaTextToNum(detailCollectormodel.OccupiedAreaText)
	usedCondoModel.OtherAreaText = detailCollectormodel.OtherAreaText
	usedCondoModel.OtherAreaNum = areaTextToNum(detailCollectormodel.OtherAreaText)
	usedCondoModel.BuiltAt, _ = time.Parse(detailCollectormodel.BuiltAt, "2025年05月")
	pref, city, town, address := utils.DivideAddress(detailCollectormodel.Address)
	usedCondoModel.AddressID, _ = utils.GetAddressID(pref, city, town)
	usedCondoModel.Address = address
	usedCondoModel.SalesSchedule = detailCollectormodel.SalesSchedule
	usedCondoModel.EventInfo = detailCollectormodel.EventInfo
	usedCondoModel.MostPriceRange = detailCollectormodel.MostPriceRange
	usedCondoModel.MaintenanceFee = detailCollectormodel.MaintenanceFee
	usedCondoModel.RepairReserveFund = detailCollectormodel.RepairReserveFund
	usedCondoModel.IntialRepairReserveFund = detailCollectormodel.IntialRepairReserveFund
	usedCondoModel.AdditionalCosts = detailCollectormodel.AdditionalCosts
	usedCondoModel.Floor = detailCollectormodel.Floor
	usedCondoModel.Direction = detailCollectormodel.Direction
	usedCondoModel.EnergyEfficiency = detailCollectormodel.EnergyEfficiency
	usedCondoModel.InsulationPerformance = detailCollectormodel.InsulationPerformance
	usedCondoModel.EstimatedUtilityCost = detailCollectormodel.EstimatedUtilityCost
	usedCondoModel.Reform = detailCollectormodel.Reform
	usedCondoModel.Structure = detailCollectormodel.Structure
	usedCondoModel.SiteArea = detailCollectormodel.SiteArea
	usedCondoModel.LandRight = detailCollectormodel.LandRight
	usedCondoModel.Zoning = detailCollectormodel.Zoning
	usedCondoModel.Parking = detailCollectormodel.Parking
	usedCondoModel.Contractor = detailCollectormodel.Contractor
	usedCondoModel.ListedFlag = true

	return usedCondoModel
}

func priceTextToNum(priceText string) int {
	priceRegex := regexp.MustCompile(`(?:(\d+)億)?(?:(\d+)万円)?`)
	var priceNum int
	if matches := priceRegex.FindStringSubmatch(priceText); len(matches) == 3 {
		oku, man := 0, 0
		if matches[1] != "" {
			if okuVal, err := strconv.Atoi(matches[1]); err == nil {
				oku = okuVal * 10000 // 億 → 万円
			}
		}
		if matches[2] != "" {
			if manVal, err := strconv.Atoi(matches[2]); err == nil {
				man = manVal
			}
		}
		priceNum = oku + man
	}
	return priceNum
}

func totalUnitsTextToNum(totalUnitsText string) int {
	totalUnitsRegex := regexp.MustCompile(`(\d+)戸`)
	matches := totalUnitsRegex.FindStringSubmatch(totalUnitsText)
	if len(matches) < 2 {
		return 0 // マッチしなかった場合は 0 を返す（もしくはエラー処理）
	}
	totalUnitsNum, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0 // 数値変換エラー時も 0 を返す（もしくはエラー処理）
	}
	return totalUnitsNum
}

func areaTextToNum(occupiedAreaText string) float64 {
	var occupiedAreaNum float64
	occupiedAreaRegex := regexp.MustCompile(`(\d+(?:\.\d+)?)m2`)
	matches := occupiedAreaRegex.FindStringSubmatch(occupiedAreaText)
	if len(matches) < 2 {
		return 0 // マッチしなければ 0 を返す
	}
	occupiedAreaNum, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return 0 // 数値変換に失敗しても 0 を返す
	}
	return occupiedAreaNum
}

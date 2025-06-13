package parse

import (
	"log"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
	"github.com/okb97/realestate-scraper/constants"
	"github.com/okb97/realestate-scraper/internal/model"
)

func ParseDetailCollector(e *colly.HTMLElement) model.DetailCollector {
	var detailCollectormodel model.DetailCollector

	detailCollectormodel.Url = e.Request.URL.String()
	e.ForEach("div.section_h2", func(_ int, el *colly.HTMLElement) {
		el.ForEach("table tr", func(_ int, tr *colly.HTMLElement) {
			e.ForEach("th", func(_ int, th *colly.HTMLElement) {
				thText := strings.Join(strings.Fields(th.Text), " ")
				thText = strings.Split(thText, "ヒント")[0]
				thText = strings.ReplaceAll(thText, " ", "")
				//fmt.Println(thText)
				//fmt.Print(": ")
				td := th.DOM.Next().Text()
				td = strings.Join(strings.Fields(td), " ")
				//fmt.Println(td)
				switch thText {
				case constants.UsedCondoName:
					detailCollectormodel.UsedCondoName = td
					//fmt.Print(td)
				case constants.PriceText:
					td := strings.Split(td, "[")[0]
					temptd := strings.ReplaceAll(td, " ", "")
					detailCollectormodel.PriceText = temptd
				case constants.Layout:
					detailCollectormodel.Layout = td
				case constants.TotalUnitsText:
					detailCollectormodel.TotalUnitsText = td
				case constants.OccupiedAreaText:
					detailCollectormodel.OccupiedAreaText = td
				case constants.OtherAreaText:
					detailCollectormodel.OtherAreaText = td
				case constants.BuiltAt:
					detailCollectormodel.BuiltAt = td
					log.Printf("BuiltAtデータを検出: キー='%s', 値='%s'", thText, td)
				case constants.Address:
					reRemove := regexp.MustCompile(`\s?\[ ■周辺環境 \]`)
					temptd := reRemove.ReplaceAllString(td, "")
					detailCollectormodel.Address = temptd
				case constants.Transportation:
					reRemove := regexp.MustCompile(`\s?\[ 乗り換え案内 \]`)
					cleanedText := reRemove.ReplaceAllString(td, "")
					// 正規表現で「歩~分」の部分を抽出
					reWalk := regexp.MustCompile(`[^[]+歩\d+分[^[]*`)
					matches := reWalk.FindAllString(cleanedText, -1)
					detailCollectormodel.Transportation = matches
				case constants.SalesSchedule:
					detailCollectormodel.SalesSchedule = td
				case constants.EventInfo:
					detailCollectormodel.EventInfo = td
				case constants.MostPriceRange:
					detailCollectormodel.MostPriceRange = td
				case constants.MaintenanceFee:
					detailCollectormodel.MaintenanceFee = td
				case constants.RepairReserveFund:
					detailCollectormodel.RepairReserveFund = td
				case constants.IntialRepairReserveFund:
					detailCollectormodel.IntialRepairReserveFund = td
				case constants.AdditionalCosts:
					detailCollectormodel.AdditionalCosts = td
				case constants.Floor:
					detailCollectormodel.Floor = td
				case constants.Direction:
					detailCollectormodel.Direction = td
				case constants.EnergyEfficiency:
					detailCollectormodel.EnergyEfficiency = td
				case constants.InsulationPerformance:
					detailCollectormodel.InsulationPerformance = td
				case constants.EstimatedUtilityCost:
					detailCollectormodel.EstimatedUtilityCost = td
				case constants.Reform:
					detailCollectormodel.Reform = td
				case constants.Structure:
					detailCollectormodel.Structure = td
				case constants.SiteArea:
					detailCollectormodel.SiteArea = td
				case constants.LandRight:
					detailCollectormodel.LandRight = td
				case constants.Zoning:
					detailCollectormodel.Zoning = td
				case constants.Parking:
					detailCollectormodel.Parking = td
				case constants.Contractor:
					detailCollectormodel.Contractor = td
				}
			})
		})
	})
	//fmt.Printf("%+v\n", detailCollectormodel)
	return detailCollectormodel

}

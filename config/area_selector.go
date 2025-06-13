package config

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// AreaSelector は地域選択を管理する構造体
type AreaSelector struct {
	areas []Area
}

// NewAreaSelector は新しいAreaSelectorを作成する
func NewAreaSelector() *AreaSelector {
	return &AreaSelector{
		areas: GetScrapeAreas(),
	}
}

// ShowMenu は地域選択メニューを表示する
func (as *AreaSelector) ShowMenu() {
	fmt.Println("=== スクレイピング対象地域選択 ===")
	fmt.Println("スクレイピングを実行する地域を選択してください:")
	fmt.Println()
	
	selectableAreas := as.areas
	
	for i, area := range selectableAreas {
		urlCount := len(area.AreaURLs)
		fmt.Printf("%d. %s (%d件のURL)\n", i+1, area.Name, urlCount)
	}
	fmt.Printf("%d. 全ての地域\n", len(selectableAreas)+1)
	fmt.Printf("%d. 終了\n", len(selectableAreas)+2)
	fmt.Println()
}

// SelectArea はユーザーの選択を受け付けて、対応するURLを返す
func (as *AreaSelector) SelectArea() ([]string, string, error) {
	selectableAreas := as.areas
	
	as.ShowMenu()
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("選択してください (1-" + strconv.Itoa(len(selectableAreas)+2) + "): ")
	
	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, "", fmt.Errorf("入力エラー: %v", err)
	}
	
	input = strings.TrimSpace(input)
	choice, err := strconv.Atoi(input)
	if err != nil {
		return nil, "", fmt.Errorf("無効な入力です。数字を入力してください")
	}
	
	if choice < 1 || choice > len(selectableAreas)+2 {
		return nil, "", fmt.Errorf("選択肢の範囲外です。1-%d の数字を入力してください", len(selectableAreas)+2)
	}
	
	// 終了が選択された場合
	if choice == len(selectableAreas)+2 {
		return nil, "exit", nil
	}
	
	// 全ての地域が選択された場合
	if choice == len(selectableAreas)+1 {
		var allURLs []string
		for _, area := range selectableAreas {
			for _, areaURL := range area.AreaURLs {
				allURLs = append(allURLs, areaURL.URL)
			}
		}
		return allURLs, "全ての地域", nil
	}
	
	// 個別の地域が選択された場合
	selectedArea := selectableAreas[choice-1]
	var urls []string
	for _, areaURL := range selectedArea.AreaURLs {
		urls = append(urls, areaURL.URL)
	}
	return urls, selectedArea.Name, nil
}

// GetAvailableAreas は利用可能な地域の一覧を返す
func (as *AreaSelector) GetAvailableAreas() []Area {
	var availableAreas []Area
	for _, area := range as.areas {
		if len(area.AreaURLs) > 0 {
			availableAreas = append(availableAreas, area)
		}
	}
	return availableAreas
}

// GetAreaSummary は地域の概要情報を返す
func (as *AreaSelector) GetAreaSummary() {
	fmt.Println("=== 設定済み地域概要 ===")
	totalURLs := 0
	
	for _, area := range as.areas {
		urlCount := len(area.AreaURLs)
		fmt.Printf("- %d: %s (%d件のURL)\n", area.ID, area.Name, urlCount)
		totalURLs += urlCount
	}
	
	fmt.Printf("\n合計利用可能URL数: %d件\n", totalURLs)
	fmt.Println()
}

// ShowUsage はコマンドライン引数の使い方を表示する
func ShowUsage() {
	fmt.Println("=== スクレーパー使用方法 ===")
	fmt.Println("コマンドライン引数で実行対象を指定してください:")
	fmt.Println()
	fmt.Println("使用例:")
	fmt.Println("  go run cmd/scraper/main.go                     # インタラクティブモード")
	fmt.Println("  go run cmd/scraper/main.go -area=1             # 東京都心部")
	fmt.Println("  go run cmd/scraper/main.go -area=1,2,3         # 複数エリア")
	fmt.Println("  go run cmd/scraper/main.go -area=1:100         # 千代田区のみ")
	fmt.Println("  go run cmd/scraper/main.go -area=1:100,101,102 # 千代田区、中央区、港区")
	fmt.Println("  go run cmd/scraper/main.go -area=7:100,101     # 横浜市鶴見区・神奈川区")
	fmt.Println("  go run cmd/scraper/main.go -list               # 利用可能地域一覧")
	fmt.Println()
	
	fmt.Println("地域番号:")
	areas := GetScrapeAreas()
	for _, area := range areas {
		fmt.Printf("  %d: %s\n", area.ID, area.Name)
	}
	fmt.Println()
	
	fmt.Println("URL番号: 各地域内で100から順番")
	fmt.Println("  例: 1:100=千代田区, 1:101=中央区, 7:100=横浜市鶴見区")
}

// ParseCommandLineArgs はコマンドライン引数を解析する
func ParseCommandLineArgs() ([]string, string, error) {
	var areaFlag = flag.String("area", "", "スクレイピング対象地域/URL (例: 1, 1:100, 1,2,3)")
	var listFlag = flag.Bool("list", false, "利用可能地域一覧を表示")
	var helpFlag = flag.Bool("help", false, "使用方法を表示")
	
	flag.Parse()
	
	if *helpFlag {
		ShowUsage()
		os.Exit(0)
	}
	
	if *listFlag {
		selector := NewAreaSelector()
		selector.GetAreaSummary()
		selector.ShowDetailedList()
		os.Exit(0)
	}
	
	if *areaFlag != "" {
		urls, err := GetSpecificURLs(*areaFlag)
		if err != nil {
			return nil, "", fmt.Errorf("引数解析エラー: %v", err)
		}
		return urls, fmt.Sprintf("指定地域 (%s)", *areaFlag), nil
	}
	
	// 引数が指定されていない場合はインタラクティブモード
	return nil, "", nil
}

// ShowDetailedList は詳細な地域・URL一覧を表示する
func (as *AreaSelector) ShowDetailedList() {
	fmt.Println("=== 詳細地域・URL一覧 ===")
	
	for _, area := range as.areas {
		fmt.Printf("\n地域 %d: %s\n", area.ID, area.Name)
		for _, areaURL := range area.AreaURLs {
			fmt.Printf("  %d:%d - %s\n", area.ID, areaURL.ID, areaURL.Name)
		}
	}
	fmt.Println()
}
package config

import (
	"fmt"
	"strconv"
	"strings"
)

// AreaURL は個別のスクレイピングURLの情報を管理する構造体
type AreaURL struct {
	ID          int    // URL ID（地域内で100から順番）
	Name        string // 区市町村名
	URL         string // スクレイピング対象URL
	Description string // 詳細説明
}

// Area はスクレイピング対象エリアの情報を管理する構造体
type Area struct {
	ID       int       // 地域ID（1-10）
	Name     string    // エリア名（例：東京都心部、横浜市）
	Code     string    // エリアコード（例：tokyo-center, yokohama）
	BaseURL  string    // エリアの基本URL
	AreaURLs []AreaURL // 実際のスクレイピング対象URL一覧
}

// GetScrapeAreas はスクレイピング対象エリアの一覧を返す
func GetScrapeAreas() []Area {
	return []Area{
		{
			ID:      1,
			Name:    "東京都心部",
			Code:    "tokyo-center",
			BaseURL: "https://suumo.jp/ms/chuko/tokyo/city/",
			AreaURLs: []AreaURL{
				{ID: 100, Name: "千代田区", URL: "https://suumo.jp/ms/chuko/tokyo/sc_chiyoda/", Description: "東京都心部 - 千代田区"},
				{ID: 101, Name: "中央区", URL: "https://suumo.jp/ms/chuko/tokyo/sc_chuo/", Description: "東京都心部 - 中央区"},
				{ID: 102, Name: "港区", URL: "https://suumo.jp/ms/chuko/tokyo/sc_minato/", Description: "東京都心部 - 港区"},
				{ID: 103, Name: "新宿区", URL: "https://suumo.jp/ms/chuko/tokyo/sc_shinjuku/", Description: "東京都心部 - 新宿区"},
				{ID: 104, Name: "文京区", URL: "https://suumo.jp/ms/chuko/tokyo/sc_bunkyo/", Description: "東京都心部 - 文京区"},
				{ID: 105, Name: "渋谷区", URL: "https://suumo.jp/ms/chuko/tokyo/sc_shibuya/", Description: "東京都心部 - 渋谷区"},
			},
		},
		{
			ID:      2,
			Name:    "東京23区東部",
			Code:    "tokyo-east",
			BaseURL: "https://suumo.jp/ms/chuko/tokyo/city/",
			AreaURLs: []AreaURL{
				{ID: 100, Name: "台東区", URL: "https://suumo.jp/ms/chuko/tokyo/sc_taito/", Description: "東京23区東部 - 台東区"},
				{ID: 101, Name: "墨田区", URL: "https://suumo.jp/ms/chuko/tokyo/sc_sumida/", Description: "東京23区東部 - 墨田区"},
				{ID: 102, Name: "江東区", URL: "https://suumo.jp/ms/chuko/tokyo/sc_koto/", Description: "東京23区東部 - 江東区"},
				{ID: 103, Name: "荒川区", URL: "https://suumo.jp/ms/chuko/tokyo/sc_arakawa/", Description: "東京23区東部 - 荒川区"},
				{ID: 104, Name: "足立区", URL: "https://suumo.jp/ms/chuko/tokyo/sc_adachi/", Description: "東京23区東部 - 足立区"},
				{ID: 105, Name: "葛飾区", URL: "https://suumo.jp/ms/chuko/tokyo/sc_katsushika/", Description: "東京23区東部 - 葛飾区"},
				{ID: 106, Name: "江戸川区", URL: "https://suumo.jp/ms/chuko/tokyo/sc_edogawa/", Description: "東京23区東部 - 江戸川区"},
			},
		},
		{
			ID:      3,
			Name:    "東京23区南部",
			Code:    "tokyo-south",
			BaseURL: "https://suumo.jp/ms/chuko/tokyo/city/",
			AreaURLs: []AreaURL{
				{ID: 100, Name: "品川区", URL: "https://suumo.jp/ms/chuko/tokyo/sc_shinagawa/", Description: "東京23区南部 - 品川区"},
				{ID: 101, Name: "目黒区", URL: "https://suumo.jp/ms/chuko/tokyo/sc_meguro/", Description: "東京23区南部 - 目黒区"},
				{ID: 102, Name: "大田区", URL: "https://suumo.jp/ms/chuko/tokyo/sc_ota/", Description: "東京23区南部 - 大田区"},
				{ID: 103, Name: "世田谷区", URL: "https://suumo.jp/ms/chuko/tokyo/sc_setagaya/", Description: "東京23区南部 - 世田谷区"},
			},
		},
		{
			ID:      4,
			Name:    "東京23区西部",
			Code:    "tokyo-west",
			BaseURL: "https://suumo.jp/ms/chuko/tokyo/city/",
			AreaURLs: []AreaURL{
				{ID: 100, Name: "中野区", URL: "https://suumo.jp/ms/chuko/tokyo/sc_nakano/", Description: "東京23区西部 - 中野区"},
				{ID: 101, Name: "杉並区", URL: "https://suumo.jp/ms/chuko/tokyo/sc_suginami/", Description: "東京23区西部 - 杉並区"},
				{ID: 102, Name: "練馬区", URL: "https://suumo.jp/ms/chuko/tokyo/sc_nerima/", Description: "東京23区西部 - 練馬区"},
			},
		},
		{
			ID:      5,
			Name:    "東京23区北部",
			Code:    "tokyo-north",
			BaseURL: "https://suumo.jp/ms/chuko/tokyo/city/",
			AreaURLs: []AreaURL{
				{ID: 100, Name: "豊島区", URL: "https://suumo.jp/ms/chuko/tokyo/sc_toshima/", Description: "東京23区北部 - 豊島区"},
				{ID: 101, Name: "北区", URL: "https://suumo.jp/ms/chuko/tokyo/sc_kita/", Description: "東京23区北部 - 北区"},
				{ID: 102, Name: "板橋区", URL: "https://suumo.jp/ms/chuko/tokyo/sc_itabashi/", Description: "東京23区北部 - 板橋区"},
			},
		},
		{
			ID:      6,
			Name:    "東京都下",
			Code:    "tokyo-tama",
			BaseURL: "https://suumo.jp/ms/chuko/tokyo/city/",
			AreaURLs: []AreaURL{
				{ID: 100, Name: "八王子市", URL: "https://suumo.jp/ms/chuko/tokyo/sc_hachioji/", Description: "東京都下 - 八王子市"},
				{ID: 101, Name: "立川市", URL: "https://suumo.jp/ms/chuko/tokyo/sc_tachikawa/", Description: "東京都下 - 立川市"},
				{ID: 102, Name: "武蔵野市", URL: "https://suumo.jp/ms/chuko/tokyo/sc_musashino/", Description: "東京都下 - 武蔵野市"},
				{ID: 103, Name: "三鷹市", URL: "https://suumo.jp/ms/chuko/tokyo/sc_mitaka/", Description: "東京都下 - 三鷹市"},
				{ID: 104, Name: "青梅市", URL: "https://suumo.jp/ms/chuko/tokyo/sc_ome/", Description: "東京都下 - 青梅市"},
				{ID: 105, Name: "府中市", URL: "https://suumo.jp/ms/chuko/tokyo/sc_fuchu/", Description: "東京都下 - 府中市"},
				{ID: 106, Name: "昭島市", URL: "https://suumo.jp/ms/chuko/tokyo/sc_akishima/", Description: "東京都下 - 昭島市"},
				{ID: 107, Name: "調布市", URL: "https://suumo.jp/ms/chuko/tokyo/sc_chofu/", Description: "東京都下 - 調布市"},
				{ID: 108, Name: "町田市", URL: "https://suumo.jp/ms/chuko/tokyo/sc_machida/", Description: "東京都下 - 町田市"},
				{ID: 109, Name: "小金井市", URL: "https://suumo.jp/ms/chuko/tokyo/sc_koganei/", Description: "東京都下 - 小金井市"},
				{ID: 110, Name: "小平市", URL: "https://suumo.jp/ms/chuko/tokyo/sc_kodaira/", Description: "東京都下 - 小平市"},
				{ID: 111, Name: "日野市", URL: "https://suumo.jp/ms/chuko/tokyo/sc_hino/", Description: "東京都下 - 日野市"},
				{ID: 112, Name: "東村山市", URL: "https://suumo.jp/ms/chuko/tokyo/sc_higashimurayama/", Description: "東京都下 - 東村山市"},
				{ID: 113, Name: "国分寺市", URL: "https://suumo.jp/ms/chuko/tokyo/sc_kokubunji/", Description: "東京都下 - 国分寺市"},
				{ID: 114, Name: "国立市", URL: "https://suumo.jp/ms/chuko/tokyo/sc_kunitachi/", Description: "東京都下 - 国立市"},
				{ID: 115, Name: "福生市", URL: "https://suumo.jp/ms/chuko/tokyo/sc_fussa/", Description: "東京都下 - 福生市"},
				{ID: 116, Name: "狛江市", URL: "https://suumo.jp/ms/chuko/tokyo/sc_komae/", Description: "東京都下 - 狛江市"},
				{ID: 117, Name: "東大和市", URL: "https://suumo.jp/ms/chuko/tokyo/sc_higashiyamato/", Description: "東京都下 - 東大和市"},
				{ID: 118, Name: "清瀬市", URL: "https://suumo.jp/ms/chuko/tokyo/sc_kiyose/", Description: "東京都下 - 清瀬市"},
				{ID: 119, Name: "東久留米市", URL: "https://suumo.jp/ms/chuko/tokyo/sc_higashikurume/", Description: "東京都下 - 東久留米市"},
				{ID: 120, Name: "武蔵村山市", URL: "https://suumo.jp/ms/chuko/tokyo/sc_musashimurayama/", Description: "東京都下 - 武蔵村山市"},
				{ID: 121, Name: "多摩市", URL: "https://suumo.jp/ms/chuko/tokyo/sc_tama/", Description: "東京都下 - 多摩市"},
				{ID: 122, Name: "稲城市", URL: "https://suumo.jp/ms/chuko/tokyo/sc_inagi/", Description: "東京都下 - 稲城市"},
				{ID: 123, Name: "羽村市", URL: "https://suumo.jp/ms/chuko/tokyo/sc_hamura/", Description: "東京都下 - 羽村市"},
				{ID: 124, Name: "あきる野市", URL: "https://suumo.jp/ms/chuko/tokyo/sc_akiruno/", Description: "東京都下 - あきる野市"},
				{ID: 125, Name: "西東京市", URL: "https://suumo.jp/ms/chuko/tokyo/sc_nishitokyo/", Description: "東京都下 - 西東京市"},
				{ID: 126, Name: "瑞穂町", URL: "https://suumo.jp/ms/chuko/tokyo/sc_mizuho/", Description: "東京都下 - 瑞穂町"},
				{ID: 127, Name: "日の出町", URL: "https://suumo.jp/ms/chuko/tokyo/sc_hinode/", Description: "東京都下 - 日の出町"},
				{ID: 128, Name: "檜原村", URL: "https://suumo.jp/ms/chuko/tokyo/sc_hinohara/", Description: "東京都下 - 檜原村"},
				{ID: 129, Name: "奥多摩町", URL: "https://suumo.jp/ms/chuko/tokyo/sc_okutama/", Description: "東京都下 - 奥多摩町"},
				{ID: 130, Name: "大島町", URL: "https://suumo.jp/ms/chuko/tokyo/sc_oshima/", Description: "東京都下 - 大島町"},
				{ID: 131, Name: "利島村", URL: "https://suumo.jp/ms/chuko/tokyo/sc_toshimamura/", Description: "東京都下 - 利島村"},
				{ID: 132, Name: "新島村", URL: "https://suumo.jp/ms/chuko/tokyo/sc_niijima/", Description: "東京都下 - 新島村"},
				{ID: 133, Name: "神津島村", URL: "https://suumo.jp/ms/chuko/tokyo/sc_kozushima/", Description: "東京都下 - 神津島村"},
				{ID: 134, Name: "三宅村", URL: "https://suumo.jp/ms/chuko/tokyo/sc_miyake/", Description: "東京都下 - 三宅村"},
				{ID: 135, Name: "御蔵島村", URL: "https://suumo.jp/ms/chuko/tokyo/sc_mikurajima/", Description: "東京都下 - 御蔵島村"},
				{ID: 136, Name: "八丈町", URL: "https://suumo.jp/ms/chuko/tokyo/sc_hachijo/", Description: "東京都下 - 八丈町"},
				{ID: 137, Name: "青ヶ島村", URL: "https://suumo.jp/ms/chuko/tokyo/sc_aogashima/", Description: "東京都下 - 青ヶ島村"},
			},
		},
		{
			ID:      7,
			Name:    "横浜市",
			Code:    "yokohama",
			BaseURL: "https://suumo.jp/ms/chuko/kanagawa/city/",
			AreaURLs: []AreaURL{
				{ID: 100, Name: "横浜市鶴見区", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_yokohamashitsurumi/", Description: "横浜市 - 鶴見区"},
				{ID: 101, Name: "横浜市神奈川区", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_yokohamashikanagawa/", Description: "横浜市 - 神奈川区"},
				{ID: 102, Name: "横浜市西区", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_yokohamashinishi/", Description: "横浜市 - 西区"},
				{ID: 103, Name: "横浜市中区", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_yokohamashinaka/", Description: "横浜市 - 中区"},
				{ID: 104, Name: "横浜市南区", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_yokohamashiminami/", Description: "横浜市 - 南区"},
				{ID: 105, Name: "横浜市保土ケ谷区", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_yokohamashihodogaya/", Description: "横浜市 - 保土ケ谷区"},
				{ID: 106, Name: "横浜市磯子区", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_yokohamashiisogo/", Description: "横浜市 - 磯子区"},
				{ID: 107, Name: "横浜市金沢区", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_yokohamashikanazawa/", Description: "横浜市 - 金沢区"},
				{ID: 108, Name: "横浜市港北区", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_yokohamashikohoku/", Description: "横浜市 - 港北区"},
				{ID: 109, Name: "横浜市戸塚区", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_yokohamashitotsuka/", Description: "横浜市 - 戸塚区"},
				{ID: 110, Name: "横浜市港南区", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_yokohamashikonan/", Description: "横浜市 - 港南区"},
				{ID: 111, Name: "横浜市旭区", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_yokohamashiasahi/", Description: "横浜市 - 旭区"},
				{ID: 112, Name: "横浜市緑区", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_yokohamashimidori/", Description: "横浜市 - 緑区"},
				{ID: 113, Name: "横浜市瀬谷区", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_yokohamashiseya/", Description: "横浜市 - 瀬谷区"},
				{ID: 114, Name: "横浜市栄区", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_yokohamashisakae/", Description: "横浜市 - 栄区"},
				{ID: 115, Name: "横浜市泉区", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_yokohamashiizumi/", Description: "横浜市 - 泉区"},
				{ID: 116, Name: "横浜市青葉区", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_yokohamashiaoba/", Description: "横浜市 - 青葉区"},
				{ID: 117, Name: "横浜市都筑区", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_yokohamashitsuzuki/", Description: "横浜市 - 都筑区"},
			},
		},
		{
			ID:      8,
			Name:    "川崎市",
			Code:    "kawasaki",
			BaseURL: "https://suumo.jp/ms/chuko/kanagawa/city/",
			AreaURLs: []AreaURL{
				{ID: 100, Name: "川崎市川崎区", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_kawasakishikawasaki/", Description: "川崎市 - 川崎区"},
				{ID: 101, Name: "川崎市幸区", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_kawasakishisaiwai/", Description: "川崎市 - 幸区"},
				{ID: 102, Name: "川崎市中原区", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_kawasakishinakahara/", Description: "川崎市 - 中原区"},
				{ID: 103, Name: "川崎市高津区", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_kawasakishitakatsu/", Description: "川崎市 - 高津区"},
				{ID: 104, Name: "川崎市多摩区", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_kawasakishitama/", Description: "川崎市 - 多摩区"},
				{ID: 105, Name: "川崎市宮前区", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_kawasakishimiyamae/", Description: "川崎市 - 宮前区"},
				{ID: 106, Name: "川崎市麻生区", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_kawasakishiasao/", Description: "川崎市 - 麻生区"},
			},
		},
		{
			ID:      9,
			Name:    "相模原市",
			Code:    "sagamihara",
			BaseURL: "https://suumo.jp/ms/chuko/kanagawa/city/",
			AreaURLs: []AreaURL{
				{ID: 100, Name: "相模原市緑区", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_sagamiharashimidori/", Description: "相模原市 - 緑区"},
				{ID: 101, Name: "相模原市中央区", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_sagamiharashichuo/", Description: "相模原市 - 中央区"},
				{ID: 102, Name: "相模原市南区", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_sagamiharashiminami/", Description: "相模原市 - 南区"},
			},
		},
		{
			ID:      10,
			Name:    "神奈川県その他",
			Code:    "kanagawa-others",
			BaseURL: "https://suumo.jp/ms/chuko/kanagawa/city/",
			AreaURLs: []AreaURL{
				{ID: 100, Name: "横須賀市", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_yokosuka/", Description: "神奈川県その他 - 横須賀市"},
				{ID: 101, Name: "平塚市", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_hiratsuka/", Description: "神奈川県その他 - 平塚市"},
				{ID: 102, Name: "鎌倉市", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_kamakura/", Description: "神奈川県その他 - 鎌倉市"},
				{ID: 103, Name: "藤沢市", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_fujisawa/", Description: "神奈川県その他 - 藤沢市"},
				{ID: 104, Name: "小田原市", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_odawara/", Description: "神奈川県その他 - 小田原市"},
				{ID: 105, Name: "茅ヶ崎市", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_chigasaki/", Description: "神奈川県その他 - 茅ヶ崎市"},
				{ID: 106, Name: "逗子市", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_zushi/", Description: "神奈川県その他 - 逗子市"},
				{ID: 107, Name: "三浦市", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_miura/", Description: "神奈川県その他 - 三浦市"},
				{ID: 108, Name: "秦野市", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_hadano/", Description: "神奈川県その他 - 秦野市"},
				{ID: 109, Name: "厚木市", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_atsugi/", Description: "神奈川県その他 - 厚木市"},
				{ID: 110, Name: "大和市", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_yamato/", Description: "神奈川県その他 - 大和市"},
				{ID: 111, Name: "伊勢原市", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_isehara/", Description: "神奈川県その他 - 伊勢原市"},
				{ID: 112, Name: "海老名市", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_ebina/", Description: "神奈川県その他 - 海老名市"},
				{ID: 113, Name: "座間市", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_zama/", Description: "神奈川県その他 - 座間市"},
				{ID: 114, Name: "南足柄市", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_minamiasigara/", Description: "神奈川県その他 - 南足柄市"},
				{ID: 114, Name: "綾瀬市", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_ayase/", Description: "神奈川県その他 - 綾瀬市"},
				{ID: 115, Name: "三浦郡葉山町", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_miuragun/", Description: "神奈川県その他 - 三浦郡葉山町"},
				{ID: 116, Name: "高座郡", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_kozagun/", Description: "神奈川県その他 - 高座郡"},
				{ID: 117, Name: "中群", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_nakagun/", Description: "神奈川県その他 - 中郡"},
				{ID: 118, Name: "足柄上郡", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_ashigarakamigun/", Description: "神奈川県その他 - 足柄上郡"},
				{ID: 119, Name: "愛甲郡", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_aikogun/", Description: "神奈川県その他 - 愛甲郡"},
				{ID: 120, Name: "足柄下郡", URL: "https://suumo.jp/ms/chuko/kanagawa/sc_ashigarashimogun/", Description: "神奈川県その他 - 足柄下郡"},
			},
		},
	}
}

// GetAreaByID は指定されたIDのエリア情報を返す
func GetAreaByID(areaID int) *Area {
	areas := GetScrapeAreas()
	for _, area := range areas {
		if area.ID == areaID {
			return &area
		}
	}
	return nil
}

// GetAreaByCode は指定されたコードのエリア情報を返す
func GetAreaByCode(code string) *Area {
	areas := GetScrapeAreas()
	for _, area := range areas {
		if area.Code == code {
			return &area
		}
	}
	return nil
}

// GetAllScrapeURLs は全エリアのスクレイピング対象URLを返す
func GetAllScrapeURLs() []string {
	var urls []string
	areas := GetScrapeAreas()
	for _, area := range areas {
		for _, areaURL := range area.AreaURLs {
			urls = append(urls, areaURL.URL)
		}
	}
	return urls
}

// GetAreaScrapeURLs は指定されたエリアIDのスクレイピング対象URLを返す
func GetAreaScrapeURLs(areaID int) []string {
	area := GetAreaByID(areaID)
	if area != nil {
		var urls []string
		for _, areaURL := range area.AreaURLs {
			urls = append(urls, areaURL.URL)
		}
		return urls
	}
	return []string{}
}

// GetAreaURLByID は指定されたエリアIDとURL IDのAreaURL情報を返す
func GetAreaURLByID(areaID, urlID int) *AreaURL {
	area := GetAreaByID(areaID)
	if area != nil {
		for _, areaURL := range area.AreaURLs {
			if areaURL.ID == urlID {
				return &areaURL
			}
		}
	}
	return nil
}

// ParseAreaURLID は "areaID:urlID" 形式の文字列をパースする
func ParseAreaURLID(input string) (int, int, error) {
	parts := strings.Split(input, ":")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid format. expected 'areaID:urlID'")
	}

	areaID, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid area ID: %v", err)
	}

	urlID, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid URL ID: %v", err)
	}

	return areaID, urlID, nil
}

// GetSpecificURLs は指定された形式の文字列からURLリストを返す
// 例: "1", "1:100", "1:100,101,102", "1,2,3"
func GetSpecificURLs(input string) ([]string, error) {
	var urls []string

	// カンマで分割して各部分を処理
	parts := strings.Split(input, ",")

	for _, part := range parts {
		part = strings.TrimSpace(part)

		if strings.Contains(part, ":") {
			// "areaID:urlID" 形式の場合
			areaID, urlID, err := ParseAreaURLID(part)
			if err != nil {
				return nil, err
			}

			areaURL := GetAreaURLByID(areaID, urlID)
			if areaURL == nil {
				return nil, fmt.Errorf("area URL not found: area %d, url %d", areaID, urlID)
			}
			urls = append(urls, areaURL.URL)
		} else {
			// 単純な数字の場合（エリアIDまたはURL ID）
			id, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("invalid ID: %v", err)
			}

			// エリアIDとして解釈
			areaURLs := GetAreaScrapeURLs(id)
			if len(areaURLs) == 0 {
				return nil, fmt.Errorf("area not found: %d", id)
			}
			urls = append(urls, areaURLs...)
		}
	}

	return urls, nil
}

# テーブル設計

## UsedCondo-中古マンション
|論理名|物理名|型|
|--|--|--|
|中古マンションID|UsedCondoId|SERIAL|
|URL|Url|TEXT|
|物件名|UsedCondoName|TEXT|
|価格（文字列）|PriceText|TEXT|
|価格（数字）|PriceNum|INTEGER|
|間取り|Layout|TEXT|
|総戸数（文字列）|TotalUnitsText|TEXT|
|総戸数（数字）|TotalUnitsNum|INTEGER|
|占有面積（文字列）|OccupiedAreaText|TEXT|
|占有面積（数字）|OccupiedAreaNum|INTEGER|
|その他面積（文字列）|OtherAreaText|TEXT|
|その他面積（数字）|OtherAreaNum|INTEGER|
|所在階|floor|TEXT|
|完成時期|BuiltAt|DATE|
|住所ID|AddressId|TEXT|
|住所1|Address1|TEXT|
|住所2|Address2|TEXT|
|販売スケジュール|SalesSchedule|TEXT|
|イベント情報|EventInfo|TEXT|
|最多価格帯|MostPriceRange|TEXT|
|管理費|MaintenanceFee|INTEGER|
|修繕積立金|RepairReserveFund|INTEGER|
|修繕積立基金|IntialRepairReserveFund|INTEGER|
|諸費用|AdditionalCosts|INTEGER|
|向き|Direction|TEXT|
|エネルギー消費性能|EnergyEfficiency|TEXT|
|断熱性能|InsulationPerformance|TEXT|
|目安光熱費|EstimatedUtilityCost|TEXT|
|リフォーム|Reform|TEXT|
|構造・階建|Structure|TEXT|
|敷地の権利形態|LandRight|TEXT|
|用途地域|Zoning|TEXT|
|駐車場|Parking|TEXT|
|施工|Contractor|TEXT|
|掲載フラグ|ListedFlag|BOOLEAN|
|削除フラグ|DeleteFlag|BOOLEAN|
|登録日時|RegisterDateTime|TIMESTAMP|
|更新日時|UpdateDateTime|TIMESTAMP|
|登録機能|RegisterFunction|TEXT|
|更新機能|UpdateFunction|TEXT|
|プロセスID|ProcessId|TEXT|

## Address-住所
|論理名|物理名|型|
|--|--|--|
|住所ID|AddressId|INTEGER|
|都道府県名|PrefectureName|TEXT|
|市区町村名|Municipalities|TEXT|
|住所|AddressName|TEXT|
|削除フラグ|DeleteFlag|BOOLEAN|
|登録日時|RegisterDateTime|TIMESTAMP|
|更新日時|UpdateDateTime|TIMESTAMP|
|登録機能|RegisterFunction|TEXT|
|更新機能|UpdateFunction|TEXT|
|プロセスID|ProcessId|TEXT|

## RailwayCompany-鉄道会社
|論理名|物理名|型|
|--|--|--|
|鉄道会社ID|RailwayCompanyId|TEXT|
|鉄道会社名|RailwayCompanyName|TEXT|
|削除フラグ|DeleteFlag|BOOLEAN|
|登録日時|RegisterDateTime|TIMESTAMP|
|更新日時|UpdateDateTime|TIMESTAMP|
|登録機能|RegisterFunction|TEXT|
|更新機能|UpdateFunction|TEXT|
|プロセスID|ProcessId|TEXT|

## Station-駅
|論理名|物理名|型|
|--|--|--|
|駅ID|StationId|TEXT|
|駅名|StationName|TEXT|
|鉄道会社ID|RailwayCompanyId|TEXT|
|住所ID|AddressId|INTEGER|
|削除フラグ|DeleteFlag|BOOLEAN|
|登録日時|RegisterDateTime|TIMESTAMP|
|更新日時|UpdateDateTime|TIMESTAMP|
|登録機能|RegisterFunction|TEXT|
|更新機能|UpdateFunction|TEXT|
|プロセスID|ProcessId|TEXT|

## Used_Condos-Stations-中古マンション-駅
|論理名|物理名|型|
|--|--|--|
|ID|Id|TEXT|
|中古マンションID|UsedCondoId|SERIAL|
|駅ID|StationId|TEXT|
|徒歩分数|WalkingMinutes|INTEGER|
|削除フラグ|DeleteFlag|BOOLEAN|
|登録日時|RegisterDateTime|TIMESTAMP|
|更新日時|UpdateDateTime|TIMESTAMP|
|登録機能|RegisterFunction|TEXT|
|更新機能|UpdateFunction|TEXT|
|プロセスID|ProcessId|TEXT|

## BusCompany-バス会社
|論理名|物理名|型|
|--|--|--|
|バス会社ID|BusCompanyId|TEXT|
|バス会社名|BusCompanyName|TEXT|
|削除フラグ|DeleteFlag|BOOLEAN|
|登録日時|RegisterDateTime|TIMESTAMP|
|更新日時|UpdateDateTime|TIMESTAMP|
|登録機能|RegisterFunction|TEXT|
|更新機能|UpdateFunction|TEXT|
|プロセスID|ProcessId|TEXT|

## BusStop-バス停
|論理名|物理名|型|
|--|--|--|
|バス停ID|BusStopId|TEXT|
|バス停名|BusStopName|TEXT|
|バス会社ID|BusCompanyId|TEXT|
|住所ID|AddressId|INTEGER|
|削除フラグ|DeleteFlag|BOOLEAN|
|登録日時|RegisterDateTime|TIMESTAMP|
|更新日時|UpdateDateTime|TIMESTAMP|
|登録機能|RegisterFunction|TEXT|
|更新機能|UpdateFunction|TEXT|
|プロセスID|ProcessId|TEXT|

## Used_Condos-BusStop-中古マンション-バス停
|論理名|物理名|型|
|--|--|--|
|ID|Id|TEXT|
|中古マンションID|UsedCondoId|SERIAL|
|バス停ID|BusStopId|TEXT|
|徒歩分数|WalkingMinutes|INTEGER|
|削除フラグ|DeleteFlag|BOOLEAN|
|登録日時|RegisterDateTime|TIMESTAMP|
|更新日時|UpdateDateTime|TIMESTAMP|
|登録機能|RegisterFunction|TEXT|
|更新機能|UpdateFunction|TEXT|
|プロセスID|ProcessId|TEXT|
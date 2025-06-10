package utils

import (
	"database/sql"
	"log"
	"os"
	"strings"
)

type Address struct {
	Prefecture string
	City       string
	Town       string
}

// 注釈削除用のキーワード
var annotationKeywords = []string{
	"（次の", // 全角の始まりに注意（他にも "（" だけでも良いが過剰除去になることも）
	"以下に掲載がない場合",
	"（地番）",
	"（無番地）",
}

func loadAddresses(db *sql.DB) ([]Address, error) {
	rows, err := db.Query(`SELECT prefecture, city, town FROM address ORDER BY address_id ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var addresses []Address
	for rows.Next() {
		var a Address
		if err := rows.Scan(&a.Prefecture, &a.City, &a.Town); err != nil {
			return nil, err
		}
		addresses = append(addresses, a)
	}

	return addresses, nil
}

// 注釈削除関数
func cleanAddressTail(s string) string {
	for _, keyword := range annotationKeywords {
		if idx := strings.Index(s, keyword); idx != -1 {
			return s[:idx]
		}
	}
	return s
}

// input に前方一致する最長の住所を見つける
func matchAddress(input string, addresses []Address) (string, string, string) {
	var (
		matched Address
		maxLen  int
	)

	for _, addr := range addresses {
		cleanTown := cleanAddressTail(addr.Town)
		full := addr.Prefecture + addr.City + cleanTown
		if strings.HasPrefix(input, full) && len(full) > maxLen {
			matched = addr
			maxLen = len(full)
		}
	}

	if maxLen > 0 {
		return matched.Prefecture, matched.City, matched.Town
	}
	return "", "", ""
}

func DivideAddress(input string) (string, string, string, string) {
	// PostgreSQLに接続
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	addresses, err := loadAddresses(db)
	if err != nil {
		log.Fatal(err)
	}

	pref, city, town := matchAddress(input, addresses)
	fullMatch := pref + city + town
	address := strings.TrimPrefix(input, fullMatch)
	address = strings.TrimSpace(address)

	return pref, city, town, address
}

func GetAddressID(prefecture, city, town string) (int, error) {
	// PostgreSQLに接続
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return -1, err
	}
	defer db.Close()

	var addressID int
	err = db.QueryRow(`
        SELECT address_id 
        FROM address 
        WHERE prefecture = $1 AND city = $2 AND town = $3
        LIMIT 1
    `, prefecture, city, town).Scan(&addressID)

	if err != nil {
		if err == sql.ErrNoRows {
			return -1, nil // 該当なしは 0 を返す
		}
		return -1, err
	}

	return addressID, nil
}

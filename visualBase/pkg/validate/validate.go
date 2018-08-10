package validate

import (
	"database/sql"
	"net/http"
)

func ValidateCurrency(currency string, db *sql.DB, w http.ResponseWriter) bool {
	var count float64
	err := db.QueryRow("SELECT SUM(inDenars) FROM currencies WHERE currency=(?)", currency).Scan(&count)

	if err != nil {
		return false
	}
	if count > 0.0 {
		return true
	}
	return false
}

func ValidateMerchant(Username string, db *sql.DB, w http.ResponseWriter) bool {
	var count int
	err := db.QueryRow("SELECT SUM(merchantAge) FROM merchants WHERE merchantUsername=(?)", Username).Scan(&count)
	if err != nil {
		return false
	}
	return true
}

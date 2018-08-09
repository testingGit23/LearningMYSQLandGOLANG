package opendb

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
)

//Tmpl is a global template
var Tmpl = template.Must(template.ParseGlob("../form/*"))

//Payment structure for info about one payment
type Payment struct {
	ID       int
	Merchant string
	Currency string
	Amount   float64
	Date     string
	Total    float64
}
type Merchant struct {
	Username  string
	Email     string
	Country   string
	Age       int
	Firstname string
	Lastname  string
}

type TypeCurrency struct {
	Currency string
	InDenars float64
}

//DbDetails structure for details about the database that is opened
type DbDetails struct {
	Host, Port, User, Password, Name string
}

//OpenDB opens the database
func OpenDB(databaseInfo []string) (db *sql.DB, e error, detailsAboutDB DbDetails) {

	detailsAboutDB = DbDetails{Host: "localhost", Port: "3306", User: databaseInfo[1], Password: databaseInfo[2], Name: databaseInfo[0]}

	db, err := sql.Open("mysql", databaseInfo[1]+":"+databaseInfo[2]+"@/"+databaseInfo[0])

	if err != nil {
		fmt.Println("Can't open " + databaseInfo[0])
		return nil, err, DbDetails{"", "", "", "", ""}
	}

	return db, nil, detailsAboutDB
}

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
	var count string
	err := db.QueryRow("SELECT * FROM merchants WHERE merchantUsername=(?)", Username).Scan(&count)
	if err != nil {
		return false
	}
	return true
}

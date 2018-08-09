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
func OpenDB(databaseInfo *[]string) (db *sql.DB, e error, detailsAboutDB DbDetails) {
	details := *databaseInfo
	detailsAboutDB = DbDetails{Host: "localhost", Port: "3306", User: details[1], Password: details[2], Name: details[0]}

	db, err := sql.Open("mysql", details[1]+":"+details[2]+"@/"+details[0])

	if err != nil {
		fmt.Println("Can't open " + details[0])
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
	var count int
	err := db.QueryRow("SELECT SUM(merchantAge) FROM merchants WHERE merchantUsername=(?)", Username).Scan(&count)
	if err != nil {
		return false
	}
	return true
}

func OpenDatabase(detailsAboutDB *[]string) func(w http.ResponseWriter, r *http.Request) {
	var details []string
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			database := r.FormValue("database")
			username := r.FormValue("username")
			password := r.FormValue("password")
			details = append(details, database)
			details = append(details, username)

			details = append(details, password)
			detailsAboutDB = &details
		}
		http.Redirect(w, r, "/", 301)

	}
}

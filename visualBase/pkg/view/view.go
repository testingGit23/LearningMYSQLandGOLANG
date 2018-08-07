package view

import (
	"LearningMYSQLandGOLANG/visualBase/pkg/opendb"
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func selectAllCurrencies(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT currency FROM currencies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ret []string
	for rows.Next() {
		var temp string
		err = rows.Scan(&temp)
		if err != nil {
			return nil, err
		}
		ret = append(ret, temp)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func View(db *sql.DB, detailsAboutDB opendb.DbDetails, err error) func(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			opendb.Tmpl.ExecuteTemplate(w, "NoSuchDB", detailsAboutDB)
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		rows, err := db.Query("SELECT * FROM payments WHERE paymentID=(?)", id)
		if err != nil {
			opendb.Tmpl.ExecuteTemplate(w, "NoSuchDB", detailsAboutDB)

		}
		var p opendb.Payment
		scanningPaymentsTable(rows, &p, detailsAboutDB, w)

		valutes, _ := selectAllCurrencies(db)
		/*
			for _, inDenars := range valutes {
				var sumFromQuerry float64
				var pom float64
				var temp float64
				err := db.QueryRow("SELECT SUM(amount) FROM payments WHERE merchantUsername=(?) AND currency=(?)", p.Merchant, inDenars).Scan(&temp)
				if err != nil {
					temp = 0
				}
				sumFromQuerry = temp

				var InDenars float64

				InDenars = selectingInDenarsFromCurrenciesTable(&InDenars, db, inDenars, w, detailsAboutDB)

				pom = sumFromQuerry * InDenars
				p.Total = p.Total + pom

			}*/
		//fmt.Println(total)
		p.Total = calculatingTotal(valutes, db, w, detailsAboutDB, p.Merchant)

		opendb.Tmpl.ExecuteTemplate(w, "Show", p)
	}
}

func scanningPaymentsTable(rows *sql.Rows, p *opendb.Payment, detailsAboutDB opendb.DbDetails, w http.ResponseWriter) {
	//var p opendb.Payment
	for rows.Next() {
		err := rows.Scan(&p.ID, &p.Merchant, &p.Currency, &p.Amount, &p.Date)
		if err != nil {
			opendb.Tmpl.ExecuteTemplate(w, "ScanError", detailsAboutDB)
		}
	}
}

func selectingInDenarsFromCurrenciesTable(InDenars *float64, db *sql.DB, inDenars string, w http.ResponseWriter, detailsAboutDB opendb.DbDetails) float64 {
	currencyInDenars, err := db.Query("SELECT inDenars FROM currencies WHERE currency=(?)", inDenars)
	defer currencyInDenars.Close()
	if err != nil {
		opendb.Tmpl.ExecuteTemplate(w, "NoSuchDB", detailsAboutDB)
	}
	var temp float64
	for currencyInDenars.Next() {

		err = currencyInDenars.Scan(&temp)
		if err != nil {
			opendb.Tmpl.ExecuteTemplate(w, "ScanError", detailsAboutDB)
		}

	}
	return temp
}

func calculatingTotal(valutes []string, db *sql.DB, w http.ResponseWriter, detailsAboutDB opendb.DbDetails, merchant string) float64 {
	var ret float64
	for _, inDenars := range valutes {
		var sumFromQuerry float64
		var pom float64
		var temp float64

		err := db.QueryRow("SELECT SUM(amount) FROM payments WHERE merchantUsername=(?) AND currency=(?)", merchant, inDenars).Scan(&temp)
		if err != nil {
			temp = 0
		}
		sumFromQuerry = temp

		var InDenars float64

		InDenars = selectingInDenarsFromCurrenciesTable(&InDenars, db, inDenars, w, detailsAboutDB)

		pom = sumFromQuerry * InDenars
		ret = ret + pom

	}
	return ret
}

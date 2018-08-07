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
		for rows.Next() {
			err = rows.Scan(&p.ID, &p.Merchant, &p.Currency, &p.Amount, &p.Date)
			if err != nil {
				opendb.Tmpl.ExecuteTemplate(w, "ScanError", detailsAboutDB)
			}
		}

		valutes, _ := selectAllCurrencies(db)

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
			currencyInDenars, err := db.Query("SELECT inDenars FROM currencies WHERE currency=(?)", inDenars)
			if err != nil {
				opendb.Tmpl.ExecuteTemplate(w, "NoSuchDB", detailsAboutDB)
			}

			for currencyInDenars.Next() {
				//var temp float64
				err = currencyInDenars.Scan(&temp)
				if err != nil {
					opendb.Tmpl.ExecuteTemplate(w, "ScanError", detailsAboutDB)
				}
				InDenars = temp
			}
			defer currencyInDenars.Close()
			pom = sumFromQuerry * InDenars
			p.Total = p.Total + pom

		}
		//fmt.Println(total)

		opendb.Tmpl.ExecuteTemplate(w, "Show", p)
	}
}

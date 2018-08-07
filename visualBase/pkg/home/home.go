package home

import (
	"LearningMYSQLandGOLANG/visualBase/pkg/opendb"
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

//Home handler function for homepage
func Home(db *sql.DB, detailsAboutDB opendb.DbDetails, err error) func(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			opendb.Tmpl.ExecuteTemplate(w, "NoSuchDB", detailsAboutDB)
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM payments")
		if err != nil {
			opendb.Tmpl.ExecuteTemplate(w, "NoSuchDB", detailsAboutDB)
		} else {
			var allPayments []opendb.Payment
			for rows.Next() {
				var p opendb.Payment
				err = rows.Scan(&p.ID, &p.Merchant, &p.Currency, &p.Amount, &p.Date)
				if err != nil {
					opendb.Tmpl.ExecuteTemplate(w, "ScanError", detailsAboutDB)

				}
				allPayments = append(allPayments, p)
			}
			//fmt.Fprintln(w, allPayments)
			opendb.Tmpl.ExecuteTemplate(w, "Home", allPayments)
		}

	}
}

package merchants

import (
	"LearningMYSQLandGOLANG/visualBase/pkg/opendb"
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

//Home handler function for homepage
func Merchants(db *sql.DB, detailsAboutDB opendb.DbDetails, err error) func(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			opendb.Tmpl.ExecuteTemplate(w, "NoSuchDB", detailsAboutDB)
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM merchants")
		if err != nil {
			opendb.Tmpl.ExecuteTemplate(w, "NoSuchDB", detailsAboutDB)
		} else {
			var allMerchants []opendb.Merchant
			for rows.Next() {
				var m opendb.Merchant
				err = rows.Scan(&m.Username, &m.Email, &m.Country, &m.Age, &m.Firstname, &m.Lastname)
				if err != nil {
					opendb.Tmpl.ExecuteTemplate(w, "ScanError", detailsAboutDB)

				}
				allMerchants = append(allMerchants, m)
			}
			//fmt.Fprintln(w, allPayments)
			opendb.Tmpl.ExecuteTemplate(w, "Merchants", allMerchants)
		}

	}
}

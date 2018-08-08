package editmerchant

import (
	"LearningMYSQLandGOLANG/visualBase/pkg/opendb"
	"database/sql"
	"net/http"
)

func Editmerchant(db *sql.DB, detailsAboutDB opendb.DbDetails, err error) func(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			opendb.Tmpl.ExecuteTemplate(w, "NoSuchDB", detailsAboutDB)
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		usr := r.URL.Query().Get("usr")
		var m opendb.Merchant
		err := db.QueryRow("SELECT * FROM merchants WHERE merchantUsername=(?)", usr).Scan(&m.Username, &m.Email, &m.Country, &m.Age, &m.Firstname, &m.Lastname)
		if err != nil {
			opendb.Tmpl.ExecuteTemplate(w, "NoSuchDB", detailsAboutDB)
		}
		opendb.Tmpl.ExecuteTemplate(w, "Editmerchant", m)
	}
}

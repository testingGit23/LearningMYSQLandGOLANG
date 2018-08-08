package newmerchant

import (
	"LearningMYSQLandGOLANG/visualBase/pkg/opendb"
	"database/sql"
	"net/http"
)

func Newmerchant(db *sql.DB, detailsAboutDB opendb.DbDetails, err error) func(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			opendb.Tmpl.ExecuteTemplate(w, "NoSuchDB", detailsAboutDB)
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		opendb.Tmpl.ExecuteTemplate(w, "Newmerchant", nil)
	}
}

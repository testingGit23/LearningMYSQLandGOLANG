package delete

import (
	"LearningMYSQLandGOLANG/visualBase/pkg/opendb"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Delete(db *sql.DB, detailsAboutDB opendb.DbDetails, err error) func(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			opendb.Tmpl.ExecuteTemplate(w, "NoSuchDB", detailsAboutDB)
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		delForm, err := db.Prepare("DELETE FROM payments WHERE paymentID=?")
		if err != nil {
			opendb.Tmpl.ExecuteTemplate(w, "PreparedError", detailsAboutDB)

		}
		delForm.Exec(id)
		log.Println("DELETE")
		http.Redirect(w, r, "/", 301)
	}
}

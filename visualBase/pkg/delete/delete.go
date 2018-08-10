package delete

import (
	"LearningMYSQLandGOLANG/visualBase/pkg/opendb"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func DeletePayment(db *sql.DB, detailsAboutDB opendb.DbDetails, err error) func(w http.ResponseWriter, r *http.Request) {
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
		http.Redirect(w, r, "/payments", 301)
	}
}

func DeleteCurrency(db *sql.DB, detailsAboutDB opendb.DbDetails, err error) func(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			opendb.Tmpl.ExecuteTemplate(w, "NoSuchDB", detailsAboutDB)
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		curr := r.URL.Query().Get("curr")
		delForm, err := db.Prepare("DELETE FROM currencies WHERE currency=(?)")
		if err != nil {
			opendb.Tmpl.ExecuteTemplate(w, "PreparedError", detailsAboutDB)

		}
		delForm.Exec(curr)
		log.Println("DELETE")
		http.Redirect(w, r, "/currencies", 301)
	}
}
func Deletemerchant(db *sql.DB, detailsAboutDB opendb.DbDetails, err error) func(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			opendb.Tmpl.ExecuteTemplate(w, "NoSuchDB", detailsAboutDB)
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		usr := r.URL.Query().Get("usr")
		delForm, err := db.Prepare("DELETE FROM merchants WHERE merchantUsername=?")
		if err != nil {
			opendb.Tmpl.ExecuteTemplate(w, "PreparedError", detailsAboutDB)

		}
		delForm.Exec(usr)
		log.Println("DELETE")
		http.Redirect(w, r, "/merchants", 301)
	}
}

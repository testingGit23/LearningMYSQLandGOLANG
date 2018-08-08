package update

import (
	"LearningMYSQLandGOLANG/visualBase/pkg/opendb"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Update(db *sql.DB, detailsAboutDB opendb.DbDetails, err error) func(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			opendb.Tmpl.ExecuteTemplate(w, "NoSuchDB", detailsAboutDB)
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			Merchant := r.FormValue("merchant")
			Currency := r.FormValue("currencies")
			Amount := r.FormValue("amount")
			Date := r.FormValue("date")
			id := r.FormValue("uid")
			insForm, err := db.Prepare("UPDATE payments SET merchantUsername=(?), currency=(?), amount=(?), dateOfPayment=(?) WHERE paymentID=(?)")
			if err != nil {
				opendb.Tmpl.ExecuteTemplate(w, "PreparedError", detailsAboutDB)
			}
			insForm.Exec(Merchant, Currency, Amount, Date, id)
			log.Println("UPDATE: Merchant: " + Merchant + " | Currency: " + Currency + " | Amount: " + Amount + " | Date: " + Date)
		}
		http.Redirect(w, r, "/", 301)
	}
}

func UpdateCurrency(db *sql.DB, detailsAboutDB opendb.DbDetails, err error) func(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			opendb.Tmpl.ExecuteTemplate(w, "NoSuchDB", detailsAboutDB)
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			Currency := r.FormValue("curr")
			InDenars := r.FormValue("indenars")
			insForm, err := db.Prepare("UPDATE currencies SET inDenars=(?) WHERE currency=(?)")
			if err != nil {
				opendb.Tmpl.ExecuteTemplate(w, "PreparedError", detailsAboutDB)
			}
			insForm.Exec(InDenars, Currency)
			log.Println("UPDATE: currency: " + Currency + " | inDenars: " + InDenars)
		}
		http.Redirect(w, r, "/currencies", 301)
	}
}

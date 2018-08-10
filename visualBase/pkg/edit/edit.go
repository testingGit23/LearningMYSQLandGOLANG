package edit

import (
	"LearningMYSQLandGOLANG/visualBase/pkg/opendb"
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func EditPayment(db *sql.DB, detailsAboutDB opendb.DbDetails, err error) func(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			opendb.Tmpl.ExecuteTemplate(w, "NoSuchDB", detailsAboutDB)
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		var p opendb.Payment
		err := db.QueryRow("SELECT * FROM payments WHERE paymentID=(?)", id).Scan(&p.ID, &p.Merchant, &p.Currency, &p.Amount, &p.Date)
		if err != nil {
			opendb.Tmpl.ExecuteTemplate(w, "NoSuchDB", detailsAboutDB)
		}
		opendb.Tmpl.ExecuteTemplate(w, "EditPayment", p)
	}
}

func EditCurrency(db *sql.DB, detailsAboutDB opendb.DbDetails, err error) func(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			opendb.Tmpl.ExecuteTemplate(w, "NoSuchDB", detailsAboutDB)
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		curr := r.URL.Query().Get("curr")
		var c opendb.TypeCurrency
		err := db.QueryRow("SELECT * FROM currencies WHERE currency=(?)", curr).Scan(&c.Currency, &c.InDenars)
		if err != nil {
			opendb.Tmpl.ExecuteTemplate(w, "NoSuchDB", detailsAboutDB)
		}
		opendb.Tmpl.ExecuteTemplate(w, "EditCurrency", c)
	}
}

func EditMerchant(db *sql.DB, detailsAboutDB opendb.DbDetails, err error) func(w http.ResponseWriter, r *http.Request) {
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

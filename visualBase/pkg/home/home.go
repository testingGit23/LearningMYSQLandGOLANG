package home

import (
	"LearningMYSQLandGOLANG/visualBase/pkg/opendb"
	"database/sql"
	"net/http"
)

func Home(db *sql.DB, detailsAboutDB opendb.DbDetails, err error) func(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			opendb.Tmpl.ExecuteTemplate(w, "NoSuchDB", detailsAboutDB)
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		opendb.Tmpl.ExecuteTemplate(w, "Home", nil)
	}
}

func CurrenciesTable(db *sql.DB, detailsAboutDB opendb.DbDetails, err error) func(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			opendb.Tmpl.ExecuteTemplate(w, "NoSuchDB", detailsAboutDB)
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM currencies")
		if err != nil {
			opendb.Tmpl.ExecuteTemplate(w, "NoSuchDB", detailsAboutDB)
		} else {
			var allCurrencies []opendb.TypeCurrency
			for rows.Next() {
				var c opendb.TypeCurrency
				err = rows.Scan(&c.Currency, &c.InDenars)
				if err != nil {
					opendb.Tmpl.ExecuteTemplate(w, "ScanError", detailsAboutDB)

				}
				allCurrencies = append(allCurrencies, c)
			}
			//fmt.Fprintln(w, allPayments)
			opendb.Tmpl.ExecuteTemplate(w, "Currencies", allCurrencies)
		}

	}
}

func MerchantsTable(db *sql.DB, detailsAboutDB opendb.DbDetails, err error) func(w http.ResponseWriter, r *http.Request) {
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

func PaymentsTable(db *sql.DB, detailsAboutDB opendb.DbDetails, err error) func(w http.ResponseWriter, r *http.Request) {
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
			opendb.Tmpl.ExecuteTemplate(w, "Payments", allPayments)
		}

	}
}

package insert

import (
	"LearningMYSQLandGOLANG/visualBase/pkg/opendb"
	"LearningMYSQLandGOLANG/visualBase/pkg/validate"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func InsertPayment(db *sql.DB, detailsAboutDB opendb.DbDetails, err error) func(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			opendb.Tmpl.ExecuteTemplate(w, "NoSuchDB", detailsAboutDB)
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			Merchant := r.FormValue("merchant")
			Currency := r.FormValue("currency")
			Amount := r.FormValue("amount")
			Date := r.FormValue("date")
			insForm, err := db.Prepare("INSERT INTO payments(paymentID,merchantUsername, currency, amount, dateOfPayment) VALUES(?,?,?,?,?)")
			if err != nil {
				opendb.Tmpl.ExecuteTemplate(w, "PreparedError", detailsAboutDB)
			}
			insForm.Exec(0, Merchant, Currency, Amount, Date)
			log.Println("INSERT: Merchant: " + Merchant + " | Currency: " + Currency + " | Amount: " + Amount + " | Date: " + Date)
		}
		//defer db.Close()
		http.Redirect(w, r, "/payments", 301)
	}
}

func InsertCurrency(db *sql.DB, detailsAboutDB opendb.DbDetails, err error) func(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			opendb.Tmpl.ExecuteTemplate(w, "NoSuchDB", detailsAboutDB)
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			Currency := r.FormValue("currency")
			InDenars := r.FormValue("indenars")
			indenars, err := strconv.ParseFloat(InDenars, 64)
			if err != nil {
				tc := opendb.TypeCurrency{Currency, 0}
				insForm, err := db.Prepare("UPDATE currencies SET inDenars=(?) WHERE currency=(?)")
				if err != nil {
					opendb.Tmpl.ExecuteTemplate(w, "PreparedError", detailsAboutDB)
				}
				insForm.Exec(0, Currency)
				log.Println("UPDATE: currency: " + Currency + " | inDenars: " + InDenars)
				opendb.Tmpl.ExecuteTemplate(w, "WrongAmountForNewCurrency", tc)
			} else {
				tc := opendb.TypeCurrency{Currency, indenars}
				val := validate.ValidateCurrency(Currency, db, w)

				if val == true {
					opendb.Tmpl.ExecuteTemplate(w, "CurrencyExist", tc)
				} else {
					insForm, err := db.Prepare("INSERT INTO currencies (currency,inDenars) VALUES(?,?)")
					if err != nil {
						opendb.Tmpl.ExecuteTemplate(w, "PreparedError", detailsAboutDB)
					}
					insForm.Exec(Currency, indenars)
					log.Println("INSERT: Currency: " + Currency + " | inDenars: " + InDenars)
					http.Redirect(w, r, "/currencies", 301)

				}
			}
		}
		//defer db.Close()
	}
}

func Insertmerchant(db *sql.DB, detailsAboutDB opendb.DbDetails, err error) func(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			opendb.Tmpl.ExecuteTemplate(w, "NoSuchDB", detailsAboutDB)
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			Username := r.FormValue("username")
			Email := r.FormValue("Email")
			Country := r.FormValue("Country")
			Age := r.FormValue("Age")
			Firstname := r.FormValue("Firstname")
			Lastname := r.FormValue("Lastname")
			insForm, err := db.Prepare("INSERT INTO merchants (merchantUsername, merchantEmail, merchantCountry, merchantAge, firstName, lastName) VALUES(?,?,?,?,?,?)")
			if err != nil {
				opendb.Tmpl.ExecuteTemplate(w, "PreparedError", detailsAboutDB)
			}
			insForm.Exec(Username, Email, Country, Age, Firstname, Lastname)
			log.Println("INSERT: Username: " + Username + " | Email: " + Email + " | Country: " + Country + " | Age: " + Age + " | Firstname: " + Firstname + " | Lastname: " + Lastname)
		}
		//defer db.Close()
		http.Redirect(w, r, "/merchants", 301)
	}
}

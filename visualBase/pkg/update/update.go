package update

import (
	"LearningMYSQLandGOLANG/visualBase/pkg/opendb"
	"LearningMYSQLandGOLANG/visualBase/pkg/validate"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func UpdatePayment(db *sql.DB, detailsAboutDB opendb.DbDetails, err error) func(w http.ResponseWriter, r *http.Request) {
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
			ID := r.FormValue("uid")
			id, err := strconv.Atoi(ID)
			if err != nil {
				id = 0
			}
			amount, err := strconv.ParseFloat(Amount, 64)
			if err != nil {
				p := opendb.Payment{id, Merchant, Currency, 0.0, Date, 0}
				opendb.Tmpl.ExecuteTemplate(w, "WrongAmount", p)
			} else {
				p := opendb.Payment{id, Merchant, Currency, amount, Date, 0}
				validCurrency := validate.ValidateCurrency(Currency, db, w)
				validMerchant := validate.ValidateMerchant(Merchant, db, w)
				if validCurrency == true && validMerchant == true {

					insForm, err := db.Prepare("UPDATE payments SET merchantUsername=(?), currency=(?), amount=(?), dateOfPayment=(?) WHERE paymentID=(?)")
					if err != nil {
						opendb.Tmpl.ExecuteTemplate(w, "PreparedError", detailsAboutDB)
					}
					insForm.Exec(Merchant, Currency, Amount, Date, id)
					log.Println("UPDATE: Merchant: " + Merchant + " | Currency: " + Currency + " | Amount: " + Amount + " | Date: " + Date)
					http.Redirect(w, r, "/payments", 301)
				} else if validMerchant == false {
					opendb.Tmpl.ExecuteTemplate(w, "NoSuchMerchant", p)
				} else {
					opendb.Tmpl.ExecuteTemplate(w, "NoSuchCurrency", p)
				}
			}
		}

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
			_, e := strconv.ParseFloat(InDenars, 64)
			if e != nil {
				panic(e)
			}
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

func UpdateMerchant(db *sql.DB, detailsAboutDB opendb.DbDetails, err error) func(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			opendb.Tmpl.ExecuteTemplate(w, "NoSuchDB", detailsAboutDB)
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			Username := r.FormValue("usr")
			Email := r.FormValue("Email")
			Country := r.FormValue("Country")
			Age := r.FormValue("Age")
			Firstname := r.FormValue("Firstname")
			Lastname := r.FormValue("Lastname")
			insForm, err := db.Prepare("UPDATE merchants SET merchantEmail=(?), merchantCountry=(?), merchantAge=(?), firstName=(?), lastName=(?) WHERE merchantUsername=(?)")
			if err != nil {
				opendb.Tmpl.ExecuteTemplate(w, "PreparedError", detailsAboutDB)
			}
			insForm.Exec(Email, Country, Age, Firstname, Lastname, Username)
			log.Println("INSERT: Username: " + Username + " | Email: " + Email + " | Country: " + Country + " | Age: " + Age + " | Firstname: " + Firstname + " | Lastname: " + Lastname)
		}
		http.Redirect(w, r, "/merchants", 301)
	}
}

package updatemerchant

import (
	"LearningMYSQLandGOLANG/visualBase/pkg/opendb"
	"database/sql"
	"log"
	"net/http"
)

func Updatemerchant(db *sql.DB, detailsAboutDB opendb.DbDetails, err error) func(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			opendb.Tmpl.ExecuteTemplate(w, "NoSuchDB", detailsAboutDB)
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			Username := r.FormValue("usr")
			newUsername:= r.FormValue("Username")
			Email := r.FormValue("Email")
			Country := r.FormValue("Country")
			Age := r.FormValue("Age")
			Firstname := r.FormValue("Firstname")
			Lastname := r.FormValue("Lastname")
			insForm, err := db.Prepare("UPDATE merchants SET  merchantUsername=(?), merchantEmail=(?), merchantCountry=(?), merchantAge=(?), firstName=(?), lastName=(?) WHERE merchantUsername=(?)")
			if err != nil {
				opendb.Tmpl.ExecuteTemplate(w, "PreparedError", detailsAboutDB)
			}
			insForm.Exec(newUsername, Email, Country, Age, Firstname, Lastname,Username)
			log.Println("INSERT: Username: " + Username + " | Email: " + Email + " | Country: " + Country + " | Age: " + Age + " | Firstname: " + Firstname + " | Lastname: " + Lastname)
		}
		http.Redirect(w, r, "/merchants", 301)
	}
}

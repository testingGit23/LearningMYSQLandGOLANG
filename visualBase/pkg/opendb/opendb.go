package opendb

import (
	"database/sql"
	"fmt"
	"html/template"
)

//Tmpl is a global template
var Tmpl = template.Must(template.ParseGlob("../form/*"))

//User the name of the user
const (
	User     = "root"
	Password = "12345"
	DbName   = "demodb"
)

//Payment structure for info about one payment
type Payment struct {
	ID       int
	Merchant string
	Currency string
	Amount   float64
	Date     string
	Total    float64
}
type Merchant struct {
	Username  string
	Email     string
	Country   string
	Age       int
	Firstname string
	Lastname  string
}

//DbDetails structure for details about the database that is opened
type DbDetails struct {
	Host, Port, User, Password, Name string
}

//OpenDB opens the database
func OpenDB() (db *sql.DB, e error) {
	db, err := sql.Open("mysql", User+":"+Password+"@/"+DbName)

	if err != nil {
		fmt.Println("Can't open " + DbName)
		return nil, err
	}

	return db, nil
}

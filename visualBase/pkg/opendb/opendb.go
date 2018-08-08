package opendb

import (
	"database/sql"
	"fmt"
	"html/template"
)

//Tmpl is a global template
var Tmpl = template.Must(template.ParseGlob("../form/*"))

//User the name of the user
/*const (
	User     = "root"
	Password = "12345"
	DbName   = "demodb"
)*/

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

type TypeCurrency struct {
	Currency string
	InDenars float64
}

//DbDetails structure for details about the database that is opened
type DbDetails struct {
	Host, Port, User, Password, Name string
}

//OpenDB opens the database
func OpenDB(databaseInfo []string) (db *sql.DB, e error, detailsAboutDB DbDetails) {

	detailsAboutDB = DbDetails{Host: "localhost", Port: "3306", User: databaseInfo[1], Password: databaseInfo[2], Name: databaseInfo[0]}

	db, err := sql.Open("mysql", databaseInfo[1]+":"+databaseInfo[2]+"@/"+databaseInfo[0])

	if err != nil {
		fmt.Println("Can't open " + databaseInfo[0])
		return nil, err, DbDetails{"", "", "", "", ""}
	}

	return db, nil, detailsAboutDB
}

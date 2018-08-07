package main

import (
	"LearningMYSQLandGOLANG/visualBase/pkg/delete"
	"LearningMYSQLandGOLANG/visualBase/pkg/edit"
	"LearningMYSQLandGOLANG/visualBase/pkg/home"
	"LearningMYSQLandGOLANG/visualBase/pkg/insert"
	"LearningMYSQLandGOLANG/visualBase/pkg/new"
	"LearningMYSQLandGOLANG/visualBase/pkg/opendb"
	"LearningMYSQLandGOLANG/visualBase/pkg/update"
	"LearningMYSQLandGOLANG/visualBase/pkg/view"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

/*


 */

func main() {
	detailsAboutDB := opendb.DbDetails{Host: "localhost", Port: "3306", User: opendb.User, Password: opendb.Password, Name: opendb.DbName}
	db, e := opendb.OpenDB()

	defer db.Close()

	http.HandleFunc("/", home.Home(db, detailsAboutDB, e))
	http.HandleFunc("/new", new.New(db, detailsAboutDB, e))
	http.HandleFunc("/view", view.View(db, detailsAboutDB, e))
	http.HandleFunc("/edit", edit.Edit(db, detailsAboutDB, e))
	http.HandleFunc("/insert", insert.Insert(db, detailsAboutDB, e))
	http.HandleFunc("/update", update.Update(db, detailsAboutDB, e))
	http.HandleFunc("/delete", delete.Delete(db, detailsAboutDB, e))
	http.ListenAndServe(":8080", nil)
}

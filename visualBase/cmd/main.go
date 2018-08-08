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

func main() {

	db, e, detailsAboutDB := opendb.OpenDB()

	defer db.Close()

	http.HandleFunc("/", home.Home(db, detailsAboutDB, e))
	http.HandleFunc("/editmerchant", edit.EditMerchant(db, detailsAboutDB, e))
	http.HandleFunc("/updatemerchant", update.UpdateMerchant(db, detailsAboutDB, e))
	http.HandleFunc("/merchants", home.MerchantsTable(db, detailsAboutDB, e))
	http.HandleFunc("/payments", home.PaymentsTable(db, detailsAboutDB, e))
	http.HandleFunc("/newpayments", new.NewPayment(db, detailsAboutDB, e))
	http.HandleFunc("/newmerchant", new.Newmerchant(db, detailsAboutDB, e))
	http.HandleFunc("/insertmerchant", insert.Insertmerchant(db, detailsAboutDB, e))
	http.HandleFunc("/viewpayment", view.ViewPayment(db, detailsAboutDB, e))
	http.HandleFunc("/editpayment", edit.EditPayment(db, detailsAboutDB, e))
	http.HandleFunc("/insertpayment", insert.InsertPayment(db, detailsAboutDB, e))
	http.HandleFunc("/updatepayment", update.UpdatePayment(db, detailsAboutDB, e))
	http.HandleFunc("/deletepayment", delete.DeletePayment(db, detailsAboutDB, e))
	http.HandleFunc("/currencies", home.CurrenciesTable(db, detailsAboutDB, e))
	http.HandleFunc("/newcurrency", new.NewCurrency(db, detailsAboutDB, e))
	http.HandleFunc("/editcurrency", edit.EditCurrency(db, detailsAboutDB, e))
	http.HandleFunc("/insertcurrency", insert.InsertCurrency(db, detailsAboutDB, e))
	http.HandleFunc("/updatecurrency", update.UpdateCurrency(db, detailsAboutDB, e))
	http.HandleFunc("/deletecurrency", delete.DeleteCurrency(db, detailsAboutDB, e))

	http.ListenAndServe(":8080", nil)
}

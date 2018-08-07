package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const (
	//host     = "localhost"
	//port     = "3306"
	user     = "root"
	password = "12345"
	dbName   = "demodb"
)

var tmpl = template.Must(template.ParseGlob("../form/*"))

type Payment struct {
	ID       int
	Merchant string
	Currency string
	Amount   float64
	Date     string
	Total    float64
}

func selectDatabase() (db *sql.DB) {
	db, err := sql.Open("mysql",
		user+":"+password+"@/"+dbName)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(db)
	//defer db.Close()
	return db
}

func home(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM payments")
		if err != nil {
			panic(err)
		}
		var allPayments []Payment
		for rows.Next() {
			var p Payment
			err = rows.Scan(&p.ID, &p.Merchant, &p.Currency, &p.Amount, &p.Date)
			if err != nil {
				panic(err)
			}
			allPayments = append(allPayments, p)
		}
		//fmt.Fprintln(w, allPayments)
		tmpl.ExecuteTemplate(w, "Home", allPayments)
	}
}

func new(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func view(w http.ResponseWriter, r *http.Request) {
	db := selectDatabase()
	id := r.URL.Query().Get("id")
	rows, err := db.Query("SELECT * FROM payments WHERE paymentID=(?)", id)
	if err != nil {
		return
	}
	var p Payment
	for rows.Next() {
		err = rows.Scan(&p.ID, &p.Merchant, &p.Currency, &p.Amount, &p.Date)
		if err != nil {
			panic(err)
		}
	}

	valutes, _ := selectAllCurrencies(db)

	for _, inDenars := range valutes {
		var sumFromQuerry float64
		var pom float64
		var temp float64
		err := db.QueryRow("SELECT SUM(amount) FROM payments WHERE merchantUsername=(?) AND currency=(?)", p.Merchant, inDenars).Scan(&temp)
		if err != nil {
			temp = 0
		}
		sumFromQuerry = temp

		var InDenars float64
		currencyInDenars, err := db.Query("SELECT inDenars FROM currencies WHERE currency=(?)", inDenars)
		if err != nil {
			panic(err)
		}

		for currencyInDenars.Next() {
			//var temp float64
			err = currencyInDenars.Scan(&temp)
			if err != nil {
				panic(err)
			}
			InDenars = temp
		}
		err = currencyInDenars.Err()
		if err != nil {
			panic(err)
		}
		defer currencyInDenars.Close()
		pom = sumFromQuerry * InDenars
		p.Total = p.Total + pom

	}
	//fmt.Println(total)

	tmpl.ExecuteTemplate(w, "Show", p)
}

func edit(w http.ResponseWriter, r *http.Request) {
	db := selectDatabase()
	id := r.URL.Query().Get("id")
	var p Payment
	err := db.QueryRow("SELECT * FROM payments WHERE paymentID=(?)", id).Scan(&p.ID, &p.Merchant, &p.Currency, &p.Amount, &p.Date)
	if err != nil {
		panic(err)
	}
	tmpl.ExecuteTemplate(w, "Edit", p)
}

func insert(w http.ResponseWriter, r *http.Request) {
	db := selectDatabase()
	if r.Method == "POST" {
		Merchant := r.FormValue("merchant")
		Currency := r.FormValue("currency")
		Amount := r.FormValue("amount")
		Date := r.FormValue("date")
		insForm, err := db.Prepare("INSERT INTO payments(paymentID,merchantUsername, currency, amount, dateOfPayment) VALUES(?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(0, Merchant, Currency, Amount, Date)
		log.Println("INSERT: Merchant: " + Merchant + " | Currency: " + Currency + " | Amount: " + Amount + " | Date: " + Date)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func update(w http.ResponseWriter, r *http.Request) {
	db := selectDatabase()
	if r.Method == "POST" {
		Merchant := r.FormValue("merchant")
		Currency := r.FormValue("currencies")
		Amount := r.FormValue("amount")
		Date := r.FormValue("date")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE payments SET merchantUsername=(?), currency=(?), amount=(?), dateOfPayment=(?) WHERE paymentID=(?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(Merchant, Currency, Amount, Date, id)
		log.Println("UPDATE: Merchant: " + Merchant + " | Currency: " + Currency + " | Amount: " + Amount + " | Date: " + Date)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func delete(w http.ResponseWriter, r *http.Request) {
	db := selectDatabase()
	id := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM payments WHERE paymentID=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(id)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func selectAllCurrencies(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT currency FROM currencies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ret []string
	for rows.Next() {
		var temp string
		err = rows.Scan(&temp)
		if err != nil {
			return nil, err
		}
		ret = append(ret, temp)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func main() {
	db := selectDatabase()
	defer db.Close()

	http.HandleFunc("/", home(db))
	http.HandleFunc("/new", new)
	http.HandleFunc("/view", view)
	http.HandleFunc("/edit", edit)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", delete)
	http.ListenAndServe(":8080", nil)
}

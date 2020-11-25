package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var db *sql.DB
func rollHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("simple_list.html")
		if err != nil {
			log.Fatal(err)
		}
		books, err := dbGetCars()
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, books)
	}
}
func addAutoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("simple_form.html")
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		name := r.Form.Get("name")
		country := r.Form.Get("country")
		year := r.Form.Get("year")
		price := r.Form.Get("price")
		err := dbAddCar(name, country, year, price)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func searchAutoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("simple_search.html")
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, nil)
	} else {
		t, err := template.ParseFiles("simple_search.html")
		r.ParseForm()
		name := r.Form.Get("name")
		cars, err := dbSearchCar(name)
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, cars)
	}
}
func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "4747"
		fmt.Println(port)
	}
	return ":" + port
}
func main() {
	err := dbConnect()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", rollHandler)
	http.HandleFunc("/add", addAutoHandler)
	http.HandleFunc("/search", searchAutoHandler)
	log.Fatal(http.ListenAndServe(GetPort(), nil))
}

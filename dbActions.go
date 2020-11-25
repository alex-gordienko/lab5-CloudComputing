package main
import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)
type Car struct{
	Name string
	Prod string
	Price string
	Year string
}
const (
	DB_USER = "postgres"
	DB_PASSWORD = "18ebyhwb"
	DB_NAME = "lab4DB"
)
func dbConnect() error {
	var err error
	db, err = sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
	DB_USER, DB_PASSWORD, DB_NAME))
	if err != nil {
		return err
	}
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS auto (auto_name text, auto_prodCountry text, auto_price text, auto_prodYear text)");
	err != nil {
	return err
}
return nil
}
func dbAddCar(name, country, year, price string) error {
	sqlstmt := "INSERT INTO auto VALUES ($1, $2, $3, $4)"
	_, err := db.Exec(sqlstmt, name, country, price, year)
	if err != nil {
		return err
	}
	return nil
}
func dbGetCars() ([]Car, error) {
	var cars []Car
	stmt, err := db.Prepare("SELECT auto_name, auto_prodCountry, auto_price, auto_prodYear FROM auto")
	if err != nil {
		return cars, err
	}
	res, err := stmt.Query()
	if err != nil {
		return cars, err
	}
	var tempCar Car
	for res.Next() {
		err = res.Scan(&tempCar.Name, &tempCar.Prod, &tempCar.Price, &tempCar.Year)
		if err != nil {
			return cars, err
		}
		cars = append(cars, tempCar)
	}
	return cars, err
}

func dbSearchCar(name string) ([]Car, error) {
	var cars []Car
	stmt, err := db.Prepare(fmt.Sprintf("SELECT auto_name, auto_prodCountry, auto_price, auto_prodYear FROM auto WHERE auto_name='%s'", name))
	if err != nil {
		return cars, err
	}
	res, err := stmt.Query()
	if err != nil {
		return cars, err
	}
	var tempCar Car
	for res.Next() {
		err = res.Scan(&tempCar.Name, &tempCar.Prod, &tempCar.Price, &tempCar.Year)
		if err != nil {
			return cars, err
		}
		cars = append(cars, tempCar)
	}
	return cars, err
}
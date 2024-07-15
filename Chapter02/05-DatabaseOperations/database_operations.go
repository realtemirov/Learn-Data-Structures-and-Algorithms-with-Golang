package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Customer struct {
	CustomerId   int
	CustomerName string
	SSN          string
}

func GetConnection() (database *sql.DB) {
	databaseDriver := "mysql"
	databaseHost := "127.0.0.1"
	databasePort := "3306"
	databaseUser := "root"
	databasePass := "mysql"
	databaseName := "crm"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", databaseUser, databasePass, databaseHost, databasePort, databaseName)
	database, err := sql.Open(databaseDriver, dsn)
	if err != nil {
		panic(err.Error())
	}

	return database
}

func GetCustomers() []Customer {
	var database *sql.DB = GetConnection()

	var (
		err  error
		rows *sql.Rows
	)

	rows, err = database.Query("SELECT * FROM customers ORDER BY CustomerId DESC")
	if err != nil {
		panic(err.Error())
	}

	var (
		customer  = Customer{}
		customers = []Customer{}
	)

	for rows.Next() {
		var (
			customerId   int
			customerName string
			ssn          string
		)
		err = rows.Scan(
			&customerId,
			&customerName,
			&ssn,
		)
		if err != nil {
			panic(err.Error())
		}
		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.SSN = ssn
		customers = append(customers, customer)
	}
	defer database.Close()

	return customers
}

func InsertCustomer(customer Customer) {
	var (
		database *sql.DB = GetConnection()
		err      error
		insert   *sql.Stmt
	)

	insert, err = database.Prepare("INSERT INTO customers (CustomerName, SSN) VALUES (?, ?)")
	if err != nil {
		panic(err.Error())
	}

	insert.Exec(customer.CustomerName, customer.SSN)

	defer database.Close()
}

func UpdateCustomer(customer Customer) {
	var (
		database *sql.DB = GetConnection()
		err      error
		update   *sql.Stmt
	)

	update, err = database.Prepare("UPDATE customers SET CustomerName = ?, SSN = ? WHERE CustomerId = ?")
	if err != nil {
		panic(err.Error())
	}

	update.Exec(customer.CustomerName, customer.SSN, customer.CustomerId)

	defer database.Close()
}

func DeleteCustomer(csutomer Customer) {
	var (
		database *sql.DB = GetConnection()
		err      error
		delete   *sql.Stmt
	)

	delete, err = database.Prepare("DELETE FROM customers WHERE CustomerId = ?")
	if err != nil {
		panic(err.Error())
	}

	delete.Exec(csutomer.CustomerId)

	defer database.Close()
}
func main() {
	var customers []Customer = GetCustomers()
	fmt.Println("Before Delete", customers)
	var customer Customer
	customer.CustomerId = 1
	DeleteCustomer(customer)
	customers = GetCustomers()
	fmt.Println("After Delete", customers)
}

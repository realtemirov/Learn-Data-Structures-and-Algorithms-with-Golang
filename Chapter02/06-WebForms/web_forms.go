package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

var template_html = template.Must(template.ParseGlob("./Chapter02/06-WebForms/templates/*"))

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

func GetCustomerById(customerId int) Customer {
	var (
		database *sql.DB = GetConnection()
		err      error
		rows     *sql.Rows
	)

	rows, err = database.Query("SELECT * FROM customers WHERE CustomerId = ?", customerId)
	if err != nil {
		panic(err.Error())
	}

	var customer Customer = Customer{}

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
	}

	defer database.Close()
	return customer
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

func Home(writer http.ResponseWriter, request *http.Request) {
	var customers []Customer
	customers = GetCustomers()
	log.Println(customers)
	template_html.ExecuteTemplate(writer, "Home", customers)

}

func Create(writer http.ResponseWriter, request *http.Request) {
	template_html.ExecuteTemplate(writer, "Create", nil)
}

func Insert(writer http.ResponseWriter, request *http.Request) {
	var (
		customer  Customer
		customers []Customer
	)

	customer.CustomerName = request.FormValue("customername")
	customer.SSN = request.FormValue("ssn")

	InsertCustomer(customer)

	customers = GetCustomers()

	template_html.ExecuteTemplate(writer, "Home", customers)
}

func Alter(writer http.ResponseWriter, request *http.Request) {
	var (
		customer      Customer
		customerId    int
		customerIdStr string
		customers     []Customer
	)

	customerIdStr = request.FormValue("id")

	fmt.Sscanf(customerIdStr, "%d", customerId)

	customer.CustomerId = customerId
	customer.CustomerName = request.FormValue("customername")
	customer.SSN = request.FormValue("ssn")

	UpdateCustomer(customer)

	customers = GetCustomers()
	template_html.ExecuteTemplate(writer, "Home", customers)
}

func Update(writer http.ResponseWriter, request *http.Request) {
	var (
		customerId    int
		customerIdStr string
		customer      Customer
	)

	customerIdStr = request.FormValue("id")

	fmt.Sscanf(customerIdStr, "%d", &customerId)
	customer = GetCustomerById(customerId)

	template_html.ExecuteTemplate(writer, "Update", customer)
}

func Delete(writer http.ResponseWriter, request *http.Request) {
	var (
		customerId    int
		customerIdStr string
		customer      Customer
		customers     []Customer
	)

	customerIdStr = request.FormValue("id")
	fmt.Sscanf(customerIdStr, "%d", &customerId)

	customer = GetCustomerById(customerId)
	DeleteCustomer(customer)

	customers = GetCustomers()

	template_html.ExecuteTemplate(writer, "Home", customers)
}

func View(writer http.ResponseWriter, request *http.Request) {
	var (
		customerId    int
		customerIdStr string
		customer      Customer
		customers     []Customer
	)

	customerIdStr = request.FormValue("id")

	fmt.Sscanf(customerIdStr, "%d", &customerId)
	customer = GetCustomerById(customerId)

	customers = GetCustomers()
	customers = append(customers, customer)

	template_html.ExecuteTemplate(writer, "View", customers)
}

func main() {
	log.Println("Server started on: http://localhost:8000")

	http.HandleFunc("/", Home)
	http.HandleFunc("/alter", Alter)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/view", View)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/delete", Delete)

	http.ListenAndServe(":8000", nil)
}

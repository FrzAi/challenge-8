package main

// import (
// 	"database/sql"
// 	"fmt"

// 	_ "github.com/lib/pq"
// )

// // database -> db_go_sql
// // schema -> grouping -> schema user, schema product, schema order
// // table -> table di schema user: user_profile, user_documents
// // index, function, etc

// // create database db_go_sql;
// // create schema db_go_sql;
// // create table db_go_sql.employees (
// // 	id serial primary key,
// // 	full_name varchar(50) not null,
// // 	email text unique not null,
// // 	age int not null,
// // 	division varchar(20) not null
// // );

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "koinworks"
// 	dbname   = "db_go_sql"
// )

// var (
// 	db  *sql.DB
// 	err error
// )

// type Employee struct {
// 	ID        int
// 	Full_name string
// 	Email     string
// 	Age       int
// 	Division  string
// }

// func main() {
// 	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

// 	db, err = sql.Open("postgres", dsn)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	err = db.Ping()
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Successfully connected to database")

// 	// CreateEmployee()
// 	// GetEmployees()
// 	UpdateEmployee()
// }

// // query INSERT, UPDATE, DELETE -> Exec
// // query SELECT -> Query/QueryRow

// func CreateEmployee() {
// 	var employee = Employee{}

// 	sqlStatement := `
// 	INSERT INTO db_go_sql.employees
// 	(
// 		full_name,
// 		email,
// 		age,
// 		division
// 	)
// 	VALUES (
// 		$1, $2, $3, $4
// 	)
// 	RETURNING *
// 	`

// 	err = db.QueryRow(sqlStatement, "Dandy Garda", "dandy2@mail.com", 17, "IT").
// 		Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)
// 	if err != nil {
// 		panic(err)
// 	}

// }

// func GetEmployees() {
// 	var result = []Employee{}
// 	query := `SELECT * FROM db_go_sql.employees ORDER BY id`

// 	rows, err := db.Query(query)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var temp = Employee{}
// 		err = rows.Scan(&temp.ID, &temp.Full_name, &temp.Email, &temp.Age, &temp.Division)
// 		if err != nil {
// 			panic(err)
// 		}

// 		result = append(result, temp)
// 	}

// 	fmt.Println("Result", result)
// }

// func UpdateEmployee() {
// 	query := `
// 	UPDATE db_go_sql.employees
// 	SET
// 		full_name = $2,
// 		email = $3,
// 		age = $4,
// 		division = $5
// 	WHERE
// 		id = $1;
// 	`

// 	res, err := db.Exec(query, 1, "Anang", "anang@mail.com", 18, "Marketing")
// 	if err != nil {
// 		panic(err)
// 	}

// 	count, err := res.RowsAffected()
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Rows affected", count)
// }

package db

import (
	"context"
	"fmt"
	"golang-app/httptypes"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host = "database"
	// host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "your-password"
	dbname   = "postgres"
)

// Employee contains the information to be stored in the database
type Employee struct {
	UserID       string `db:"userid"`
	EmailAddress string `db:"email"`
	FirstName    string `db:"first_name"`
	JobTitle     string `db:"job_title"`
	LastName     string `db:"last_name"`
	Region       string `db:"region"`
}

// MyDatabase contains database client
type MyDatabase struct {
	Client *sqlx.DB
}

// Connect to Postgres database
func OpenConnection() *MyDatabase {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(100)*time.Second)
	defer cancel()

	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbname, password, port)
	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	if err := db.PingContext(ctx); err != nil {
		panic(err)
	}
	return &MyDatabase{
		Client: db,
	}
}

// AddEmployee adds a new employee to the database
func (db *MyDatabase) AddEmployee(e *httptypes.Employees) error {
	for _, emp := range e.Employees {
		_, err := db.Client.Exec("INSERT INTO employee(userid,email,first_name,job_title,last_name,region) VALUES($1,$2,$3,$4,$5,$6);", emp.UserID, emp.EmailAddress, emp.FirstName, emp.JobTitle, emp.LastName, emp.Region)
		if err != nil {
			panic(err)
		}
	}
	return nil
}

// GetEmployees return all the employees from the database
func (db *MyDatabase) GetEmployees() []httptypes.Employee {
	rows, err := db.Client.Queryx("Select * from employee")
	if err != nil {
		panic(err)
	}
	employee := make([]httptypes.Employee, 0)
	for rows.Next() {
		row := Employee{}
		err = rows.StructScan(&row)
		if err != nil {
			panic(err)
		}
		temp := httptypes.Employee{
			UserID:       row.UserID,
			FirstName:    row.FirstName,
			LastName:     row.LastName,
			EmailAddress: row.EmailAddress,
			JobTitle:     row.JobTitle,
			Region:       row.Region,
		}
		employee = append(employee, temp)
	}
	return employee
}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v4"
	_ "github.com/lib/pq"
)

type Employee struct {
	FirstName   string    `faker:"first_name"`
	LastName    string    `faker:"last_name"`
	DateOfBirth time.Time `faker:"-"`
	PhoneNumber string    `faker:"phone_number"`
}

func main() {
	db, err := sql.Open(
		"postgres",
		"host=localhost user=root password=postgres dbname=gosql sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS employees (
			employee_id SERIAL PRIMARY KEY,
			first_name VARCHAR(1000) NOT NULL,
			last_name VARCHAR(1000) NOT NULL,
			date_of_birth DATE NOT NULL,
			phone_number VARCHAR(1000) NOT NULL
		)
	`)

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 1000; i++ {
		employee := Employee{}
		err := faker.FakeData(&employee)
		if err != nil {
			log.Fatal(err)
		}

		employee.DateOfBirth = randomDate()

		_, err = db.Exec(`
        INSERT INTO employees(first_name, last_name, date_of_birth, phone_number) 
        VALUES ($1, $2, $3, $4)
      `,
			employee.FirstName,
			employee.LastName,
			employee.DateOfBirth,
			employee.PhoneNumber,
		)

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Data inserted successfully")
}

func randomDate() time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Now().Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

package models

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" //Driver for postgres
)

// InitDB ... asdasd
func InitDB() (*sql.DB, error) {
	var err error
	db, err := sql.Open("postgres",
		"postgres://postgres:kuroko9.@localhost/mydb?sslmode=disable")
	if err != nil {
		return nil, err
	}
	//Create model for our URL service
	stmt, err := db.Prepare("CREATE TABLE WEB_URL (ID SERIAL PRIMARY KEY, URL TEXT NOT NULL);")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	res, err := stmt.Exec()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(res, "Succesfully created")
	return db, nil

}

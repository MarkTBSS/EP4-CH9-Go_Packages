package user

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var database *sql.DB

func InitDB() {
	var err error
	database, err = sql.Open("postgres", "user=postgres password=Pass1234 host=127.0.0.1 port=5432 sslmode=disable")
	if err != nil {
		log.Fatal("Connect to database error : ", err)
	}
	//defer database.Close()
	createTable := `CREATE TABLE IF NOT EXISTS users ( id SERIAL PRIMARY KEY, name TEXT, age INT );`
	_, err = database.Exec(createTable)
	if err != nil {
		log.Fatal("Can't create table : ", err)
	}
}

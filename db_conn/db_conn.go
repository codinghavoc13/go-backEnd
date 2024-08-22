package db_conn

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "control"
	dbname   = "monolithDB"
)

// var DB *sql.DB

func Connect() *sql.DB {
	// connStr := fmt.Sprintf("postgresql://%v:%v@localhost:5432/monolithDB?sslmode=disable",
	// 	user, password)
	// connStr := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)

	var db *sql.DB
	var err error
	// fmt.Println("ENV:", os.Getenv("ENV"))
	// fmt.Println("DB_HOST:", os.Getenv("DB_HOST"))
	// if os.Getenv("ENV") == "dev" {
	// connStr := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PW"), dbname)

	// db, err = sql.Open("postgres", connStr)

	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	// }
	// if os.Getenv("ENV") == "prod" {
	// connStr := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PW"), dbname)

	// db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	// if err != nil {
	// 	panic(err)

	// }
	// }
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

package connections

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/boil"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "dbname=dvdrental user=postgres password=admin host=localhost port=5432 sslmode=disable")
	if err != nil {
		return db, err
	}

	return db, nil
}

//to add in main package
func init() {
	db, err := InitDB()
	if err != nil {
		log.Fatal(err)
	}

	boil.SetDB(db)
	boil.DebugMode = true
}

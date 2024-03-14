package connectors

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func GetDBConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/imersao17")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func CloseDBConnection() {
	db := GetDBConnection()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("ERROR: %v", err)
		}
	}(db)
}

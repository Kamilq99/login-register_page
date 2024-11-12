package db_connect

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DBconnect(username string, password string, dbadress string, dbname string) (*sql.DB, error) {
	db, err := sql.Open("mysql", username+":"+password+dbadress+dbname)
	if err != nil {
		return nil, fmt.Errorf("Database Connection Not Reached", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Error Pinging Database", err)
	}
	return db, err
}

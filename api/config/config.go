package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connectdb() (*sql.DB, error) {
	// in a real setting, I would put this in the kubernetes config file.
	db, errdb := sql.Open("mysql", "root:root@tcp(localhost:3306)/auction")
	if errdb != nil {
		return nil, errdb
	}
	err := db.Ping()
	return db, err
}

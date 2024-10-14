package database

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)
	

func Connect() (*sql.DB, error) {
	stringConnection := "script:123456789@tcp(192.168.0.20:3306)/golang?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", stringConnection)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
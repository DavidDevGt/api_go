package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConectarDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "usuario:contraseña@tcp(localhost)/ticket_sys_db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

package service

import "database/sql"

var db *sql.DB

func SetupServiceLayer(dbConn *sql.DB) error {
	db = dbConn
	return nil
}

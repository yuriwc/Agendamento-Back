package services

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func ConnectToDB() *sqlx.Tx {
	db, err := sqlx.Connect(
		"mssql",
		"sqlserver://sa:senha@127.0.0.1:1433?database=Banco&connection+timeout=30",
	)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	tx := db.MustBegin()

	return tx
}

package log

import (
	"database/sql"

	database "github.com/Carry-Rao/go-db"
)

var errorDB *sql.DB

func initErrorDB(typ string, dsn string) {
	var err error
	errorDB, err = database.Open(typ, dsn)
	if err != nil {
		panic(err)
	}
	err = errorDB.Ping()
	if err != nil {
		panic(err)
	}

	_, err = errorDB.Exec(`
		CREATE TABLE IF NOT EXISTS error (
			time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			message TEXT NOT NULL,
			function TEXT NOT NULL
		);
	`)
	if err != nil {
		panic(err)
	}
	_, err = errorDB.Exec(`
		CREATE TABLE IF NOT EXISTS error_http (
			time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			url TEXT NOT NULL,
			method TEXT NOT NULL,
			status_code INT NOT NULL
		)
	`)
	if err != nil {
		panic(err)
	}
}

func Error(message string, function string) {
	_, err := errorDB.Exec(`
		INSERT INTO error (message, function) VALUES (?,?)
		`, message, function)
	if err != nil {
		panic(err)
	}
}

func ErrorHTTP(url string, method string, status_code int, header string) {
	_, err := errorDB.Exec(`
		INSERT INTO error_http (url, method, status_code, header) VALUES (?,?,?,?)
		`, url, method, status_code, header)
	if err != nil {
		panic(err)
	}
}

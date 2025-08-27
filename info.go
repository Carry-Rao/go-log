package log

import (
	"database/sql"

	database "github.com/Carry-Rao/go-db"
)

var infoDB *sql.DB

func initInfoDB(typ string, dsn string) {
	var err error
	infoDB, err = database.Open(typ, dsn)
	if err != nil {
		panic(err)
	}
	err = infoDB.Ping()
	if err != nil {
		panic(err)
	}

	_, err = infoDB.Exec(`
		CREATE TABLE IF NOT EXISTS info (
			time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			message TEXT NOT NULL,
			function TEXT NOT NULL
		);
	`)
	if err != nil {
		panic(err)
	}
	_, err = infoDB.Exec(`
		CREATE TABLE IF NOT EXISTS info_http (
			time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			url TEXT NOT NULL,
			method TEXT NOT NULL,
			status_code INT NOT NULL,
			header TEXT NOT NULL
		)
	`)
	if err != nil {
		panic(err)
	}
}

func Info(message string, function string) {
	_, err := infoDB.Exec(`
		INSERT INTO info (message, function) VALUES (?,?)
		`, message, function)
	if err != nil {
		panic(err)
	}
}

func InfoHttp(url string, method string, status_code int, header string) {
	_, err := infoDB.Exec(`
		INSERT INTO info_http (url, method, status_code, header) VALUES (?,?,?,?)
		`, url, method, status_code, header)
	if err != nil {
		panic(err)
	}
}

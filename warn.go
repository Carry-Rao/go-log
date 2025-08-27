package log

import (
	database "github.com/Carry-Rao/go-db"

	"database/sql"
)

var warnDB *sql.DB

func initWarnDB(typ string, dsn string) {
	var err error
	warnDB, err = database.Open(typ, dsn)
	if err != nil {
		panic(err)
	}
	err = warnDB.Ping()
	if err != nil {
		panic(err)
	}

	_, err = warnDB.Exec(`
		CREATE TABLE IF NOT EXISTS warn (
			time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			message TEXT NOT NULL,
			function TEXT NOT NULL
		);
	`)
	if err != nil {
		panic(err)
	}
	_, err = warnDB.Exec(`
		CREATE TABLE IF NOT EXISTS warn_http (
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

func Warn(message string, function string) {
	_, err := warnDB.Exec(`
		INSERT INTO warn (message, function) VALUES (?,?)
		`, message, function)
	if err != nil {
		panic(err)
	}
}

func WarnHTTP(url string, method string, status_code int, header string) {
	_, err := warnDB.Exec(`
		INSERT INTO warn_http (url, method, status_code, header) VALUES (?,?,?,?)
		`, url, method, status_code, header)
	if err != nil {
		panic(err)
	}
}
